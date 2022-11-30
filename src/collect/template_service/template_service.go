package collect

import (
	"bytes"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"fmt"
	"github.com/demdxx/gocast"
	text_template "text/template"
)

type TemplateService struct {
	OpUser  string
	Session interface{}
	Request interface{}
}

func (t *TemplateService) getModuleResultObj(moduleName string) config.ModuleResult {
	// 根据模块名称获取，模块对象
	module := config.GetModuleRegister(moduleName)
	return module

}
func IsPluginEnable(Tpl *text_template.Template, plugin config.Plugin) bool {
	var buf bytes.Buffer

	err := Tpl.Execute(&buf, plugin)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return gocast.ToBool(buf.String())

}
func (t *TemplateService) before(params map[string]interface{}, is_http bool) (*config.Template, *common.Result) {
	serviceName := config.GetServiceName(params)
	if utils.IsValueEmpty(serviceName) {
		err_msg := "请求参数【service】不存在，请检查传入参数"
		return nil, common.NotOk(err_msg)
	}

	// 根据service 名称获取配置
	cfg := config.NewTemplateService(params)
	// 生成模板
	temp := config.Template{}
	// 设置服务配置
	temp.ServiceConfig = cfg.GetData().(config.ServiceConfig)
	// 设置全局路由配置
	temp.RouterAllConfig = config.GetLocalRouter()
	// 设置参数
	temp.SetParams(params)
	// todo: 这里只示例了一个用户ID
	// 设置操作用户,需要将模块的变量，赋值给temp,比如session,event_id，http 的请求对象，http 的请求头
	// 内部的服务调用，也是如此，比如template 生成了一个事件ID,后面服务沿用这个事件ID,直到服务就结束
	temp.OpUser = t.OpUser
	if temp.Log {
		msg := "【" + temp.OpUser + "】访问:" + serviceName
		temp.LogData(msg)
	}
	var loader config.BeforeLoader
	for _, plugin := range temp.GetBeforePlugins() {
		//插件是否启用
		if !IsPluginEnable(plugin.EnableTpl, plugin) {
			continue
		}
		pluginResult := config.CallPluginFunc(&loader, plugin, temp)
		return nil, pluginResult
	}

	return &temp, common.Ok(&temp, "成功")
}

func (t *TemplateService) execute(temp *config.Template) *common.Result {
	// 调用模块结果
	result := t.getModuleResultObj(temp.Module)
	params := temp.GetParams()
	if result == nil {
		err_msg := "module【" + temp.Module + "】模块不存在，请检查配置"
		temp.LogErr(err_msg)
		temp.LogErr(params)
		return common.NotOk(err_msg)
	}
	data, _ := result.Result(temp)
	return data
}
func (t *TemplateService) after() *common.Result {
	return nil
}
func (t *TemplateService) Result(params map[string]interface{}, isHttp bool) *common.Result {

	// 执行处理前
	temp, beforeResult := t.before(params, isHttp)
	if !beforeResult.Success {
		return beforeResult
	}
	//// 调用模块结果
	//result := t.getModuleResultObj(temp.Module)
	//if result == nil {
	//	err_msg := "module【" + temp.Module + "】模块不存在，请检查配置"
	//	temp.LogErr(err_msg)
	//	temp.LogErr(params)
	//	return common.NotOk(err_msg)
	//}
	//// 运行结果
	//data, _ := result.Result(temp)
	//return data
	data := t.execute(temp)
	return data
}

/*
插件是实现方式后续支持
	m, err := plugin.Open("sql_service.so")
	if err != nil {
		return common.NotOk(err.Error())
	}
	r, err := m.Lookup("SqlServerPlugin")
	if err != nil {
		msg := err.Error()
		return common.NotOk(msg)
	}
	result := r.(ModuleResult)
	data, _ :=result.Result(&temp)

	if result, ok := r.(ModuleResult); ok {
		data, _ := result.Result(&temp)
		return data
	}
*/
