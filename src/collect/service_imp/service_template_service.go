package collect

import (
	"bytes"
	common "github.com/SelfDown/collect/src/collect/common"
	"github.com/SelfDown/collect/src/collect/config"
	startup "github.com/SelfDown/collect/src/collect/startup"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"reflect"
	"time"

	"fmt"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"github.com/gin-contrib/sessions"
	"github.com/robfig/cron/v3"
	text_template "text/template"
)

type TemplateService struct {
	OpUser           string
	session          *sessions.Session // 设置session
	Request          interface{}
	IsFileResponse   bool
	ResponseFilePath string
	ResponseFileName string
	File             multipart.File // 单个上传文件
	thirdData        map[string]interface{}
}

/*
* 数据库表
 */
var _model DatabaseModel

type DatabaseModel interface {
	GetModel(tableName string) interface{}
	CloneModel(tableName string) interface{}
	GetPrimaryKey(tableName string) []string
}

// 扩展模块设置
func (s *TemplateService) SetThirdData(key string, value interface{}) {
	if s.thirdData == nil {
		s.thirdData = make(map[string]interface{})
	}
	s.thirdData[key] = value
}
func (s *TemplateService) HasThirdData(key string) bool {
	_, has := s.thirdData[key]
	return has
}

//扩展模块获取
func (s *TemplateService) GetThirdData(key string) interface{} {
	return s.thirdData[key]
}

//扩展模块获取
func (s *TemplateService) RemoveThirdData(key string) {
	delete(s.thirdData, key)
}

// SetDatabaseModel 设置数据库表
func SetDatabaseModel(m DatabaseModel) {
	_model = m
}
func (*TemplateService) GetModel(tableName string) interface{} {
	return _model.GetModel(tableName)
}
func (*TemplateService) CloneModel(tableName string) interface{} {
	return _model.CloneModel(tableName)
}
func (*TemplateService) GetPrimaryKey(tableName string) []string {
	return _model.GetPrimaryKey(tableName)
}

func handlerAmis(result *common.Result) {
	// 将结果转成支持amis形式
	if utils.GetAppKey("to_amis") == "true" {
		if result.Success { // 将status 设置成数字
			result.Status = 0
		} else {
			result.Status = -1
		}
		dataResult := result.GetData()
		// 处理data
		if _, ok := dataResult.(map[string]interface{}); !ok && result.Success {
			rData := make(map[string]interface{})
			rData["data"] = dataResult
			// 只有查询返回list,list 才有count
			if _, isList := dataResult.([]map[string]interface{}); isList {
				rData["count"] = result.GetCount()
			}
			result.Data = rData
		}

	}
}

// RunScheduleService 添加定时任务
func RunScheduleService() []*collect.ServiceConfig {
	services := collect.GetLocalRouter().GetRegisterServices()
	scheduleService := make([]*collect.ServiceConfig, 0)
	params := make(map[string]interface{})
	c := cron.New()
	i := 0
	for _, service := range services {
		if !utils.IsValueEmpty(service.Schedule.Enable) && !utils.IsValueEmpty(service.Schedule.ScheduleSpec) {
			run := utils.RenderTplBool(service.Schedule.EnableTpl, params)
			if run {
				scheduleService = append(scheduleService, service)
				i++
				paramService := make(map[string]interface{})
				paramService["service"] = service.Service
				fmt.Println(utils.Strval(i) + ".添加定时任务[" + service.Service + "]" + service.Schedule.ScheduleSpec)
				c.AddFunc(service.Schedule.ScheduleSpec, func() {
					ts := TemplateService{OpUser: "schedule"}
					ts.ResultInner(paramService)
				})
			}
		}

	}
	if len(scheduleService) > 0 {
		c.Start()
	}
	return scheduleService
}

// RunStartupService 添加启动服务
func RunStartupService() []*collect.ServiceConfig {
	services := collect.GetLocalRouter().GetRegisterServices()
	startupService := make([]*collect.ServiceConfig, 0)
	i := 0
	for _, service := range services {
		if service.RunStartup {
			startupService = append(startupService, service)
			i++
			paramService := make(map[string]interface{})
			paramService["service"] = service.Service
			fmt.Println(utils.Strval(i) + ".执行启动服务[" + service.Service + "]" + service.Schedule.ScheduleSpec)
			ts := TemplateService{OpUser: "startup"}
			ts.ResultInner(paramService)
		}

	}

	return startupService
}

func HandlerRequest(c *gin.Context) {
	s := sessions.Default(c)
	//设置参数
	params := make(map[string]interface{})
	c.Bind(&params)

	// session 中设置用户ID
	userId := s.Get("user_id")
	var opUser string
	if userId != nil {
		opUser = userId.(string)
	} else {
		opUser = ""
	}
	ts := TemplateService{OpUser: opUser}
	if c.Request.PostForm != nil {
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				params[k] = v[0]
			} else {
				params[k] = nil
			}

		}
		// 处理单个文件上传
		file, _, error := c.Request.FormFile("file")
		if error != nil {
			data := common.NotOk(error.Error())
			c.JSON(200, data)
		}
		ts.File = file
	}
	// 设置session

	ts.SetSession(&s)
	// 处理结果
	var data *common.Result
	utils.Block{
		Try: func() {
			data = ts.Result(params, true)
		},
		Catch: func(e utils.Exception) {
			dv := reflect.ValueOf(e)
			data = common.NotOk(dv.String())
		},
	}.Do()

	// 处理amis结果
	//handlerAmis(data)
	if ts.IsFileResponse {
		filename := ts.ResponseFileName
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File(ts.ResponseFilePath)
	} else {
		c.JSON(200, data)
	}
}

func init() {
	// 加载配置
	// 加载系统插件，主要加载count、file_data,
	routerAll := startup.LoadSystemServices()
	//获取启动注册的服务列表，然后路由设置注册服务
	SetRegisterList(&routerAll, GetRegisterList())
	//设置服务
	collect.SetLocalRouter(routerAll)

}
func (t *TemplateService) SetSession(session *sessions.Session) {

	t.session = session
}

// GetSession 获取session
func (t *TemplateService) GetSession() *sessions.Session {
	return t.session
}

func (t *TemplateService) getModuleResultObj(moduleName string) ModuleResult {
	// 根据模块名称获取，模块对象
	module := GetModuleRegister(moduleName)
	return module

}
func IsPluginEnable(Tpl *text_template.Template, plugin collect.Plugin) bool {
	var buf bytes.Buffer

	err := Tpl.Execute(&buf, plugin)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return gocast.ToBool(buf.String())

}
func (t *TemplateService) before(params map[string]interface{}, isHttp bool) (*collect.Template, *common.Result) {
	serviceName := collect.GetServiceName(params)
	if utils.IsValueEmpty(serviceName) {
		errMsg := "请求参数【service】不存在，请检查传入参数"
		return nil, common.NotOk(errMsg)
	}

	// 根据service 名称获取配置
	cfg := collect.NewTemplateService(params)
	if !cfg.Success {
		return nil, cfg
	}
	// 生成模板
	temp := collect.Template{}
	// 设置服务配置
	temp.ServiceConfig = cfg.GetData().(collect.ServiceConfig)
	// 设置全局路由配置
	temp.RouterAllConfig = collect.GetLocalRouter()

	// 设置参数
	temp.SetParams(params)
	// todo: 这里只示例了一个用户ID
	// 设置操作用户,需要将模块的变量，赋值给temp,比如session,event_id，http 的请求对象，http 的请求头
	// 内部的服务调用，也是如此，比如template 生成了一个事件ID,后面服务沿用这个事件ID,直到服务就结束
	temp.OpUser = t.OpUser
	if temp.Log {
		msg := "【" + temp.OpUser + "】访问:" + serviceName
		temp.LogData(msg)
		temp.LogData(params)
	}
	// 如果是http 请求，并且配置必须登录，如果没有用户ID,则提示必须登录
	mustLogin := true
	if temp.MustLogin != nil && *temp.MustLogin == false {
		mustLogin = false
	}
	// http 登录判断
	if isHttp && !temp.Http {
		errMsg := serviceName + "不支持http 访问"
		return nil, common.NotOk(errMsg)
	}
	// 用户登录判断
	if isHttp && mustLogin && utils.IsValueEmpty(t.OpUser) {
		errMsg := "请登录！！！"
		return nil, common.NotOk(errMsg)
	}
	var loader BeforeLoader
	for _, plugin := range temp.GetBeforePlugins() {
		//插件是否启用
		if !IsPluginEnable(plugin.EnableTpl, plugin) {
			continue
		}
		pluginResult := collect.CallPluginFunc(&loader, plugin, &temp, t)
		if pluginResult.IsFinish() {
			return &temp, pluginResult
		}

	}

	return &temp, common.Ok(&temp, "成功")
}
func ExecTime(template *collect.Template, start time.Time, method string) {
	dis := time.Now().Sub(start).Seconds()
	template.LogData("服务：" + template.GetService() + " [" + method + "]耗时：" + utils.Strval(dis) + "s")
}

func (t *TemplateService) execute(temp *collect.Template) *common.Result {
	if temp.Log {
		defer ExecTime(temp, time.Now(), temp.Module)
	}
	// 调用模块结果
	result := t.getModuleResultObj(temp.Module)
	params := temp.GetParams()
	if result == nil {
		errMsg := "module【" + temp.Module + "】模块不存在，请检查配置"
		temp.LogErr(errMsg)
		temp.LogErr(params)
		return common.NotOk(errMsg)
	}
	data := result.Result(temp, t)

	return data
}
func (t *TemplateService) after(temp *collect.Template) *common.Result {

	var loader AfterLoader
	for _, plugin := range temp.GetAfterPlugins() {
		//插件是否启用
		if !IsPluginEnable(plugin.EnableTpl, plugin) {
			continue
		}
		pluginResult := collect.CallPluginFunc(&loader, plugin, temp, t)
		if !pluginResult.Success {
			return pluginResult
		}

	}

	return common.Ok(&temp, "成功")
}
func (t *TemplateService) ResultInner(params map[string]interface{}) *common.Result {
	var data *common.Result
	utils.Block{
		Try: func() {
			copyMap := utils.Copy(params).(map[string]interface{})
			data = t.Result(copyMap, false)
		},
		Catch: func(exception utils.Exception) {
			data = common.NotOk(utils.Strval(exception))
		},
	}.Do()
	return data
}
func (t *TemplateService) Result(params map[string]interface{}, isHttp bool) *common.Result {

	// 执行处理前

	temp, beforeResult := t.before(params, isHttp)
	//如果失败了，或则结束直接返回
	if beforeResult.IsFinish() {
		return beforeResult
	}
	// 处理中
	data := t.execute(temp)
	// 设置结果
	// 如果参数处理中没有设置过结果，则设置模块处理的结果 ，优先参数返回的结果，
	// 一般是空模块参数中配置了结果，如果实在参数中配置了结果，模块中也需要结果，那么参数中的结果不能配置
	// 最好result_handler 配置结果
	if !temp.HasResult() {
		temp.SetResult(data)
	}

	if !data.Success {
		return data
	}
	afterResult := t.after(temp)
	if !afterResult.Success {
		return afterResult
	}
	// 获取结果
	result := temp.GetResult()

	return result
}
