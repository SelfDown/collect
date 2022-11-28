package collect

import (
	utils "collect.mod/src/collect/utils"
	"github.com/demdxx/gocast"
)

type BeforeLoader struct {
}

func (t *BeforeLoader) HandlerReqParam(config Plugin, template Template, routerAll *RouterAll) {
	paramConfig := template.Params
	// 处理默认值
	for name, config := range paramConfig {
		//如果没有配置default返回
		if utils.IsValueEmpty(config.Default) {
			continue
		}
		//如果有值则跳过
		if !utils.IsEmpty(name, template.param_pool) {
			continue
		}
		// 根据名称设置默认值
		template.param_pool[name] = config.Default
	}
	//处理template
	for name, config := range paramConfig {
		if utils.IsValueEmpty(config.Template) {
			continue
		}
		// 模板渲染值
		value := utils.RenderTplExec(config.TemplateTpl, template.param_pool, config.Exec)
		// 模板变量赋值
		template.param_pool[name] = value
	}
	//处理数据类型
	for name, config := range paramConfig {

		if utils.IsValueEmpty(config.Type) {
			continue
		}
		// 模板渲染值
		value := template.param_pool[name]
		if utils.IsValueEmpty(value) {
			continue
		}
		switch config.Type {
		case "int":
			value = gocast.ToInt(value)
			break
		case "bool":
			value = gocast.ToBool(value)
			break
		case "float":
			value = gocast.ToFloat(value)
			break
		}
		// 重新设置值
		template.param_pool[name] = value
	}

}
