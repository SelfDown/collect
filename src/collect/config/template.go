package collect

import (
	common "collect/src/collect/common"
	"encoding/json"
	"log"

	// service_config "test.mod/src/collect/service_config"
	utils "collect/src/collect/utils"

	"gopkg.in/yaml.v2"
)

type Template struct {
	Service   string                 // 服务名称
	OpUser    string                 // 操作用户
	EventId   string                 // 事件ID
	paramPool map[string]interface{} //请求参数池
	ServiceConfig
	RouterAllConfig *RouterAll             // 总服务路由
	count           int64                  // 将模板的里面计算的count 存储到模板中
	result          *common.Result         // 模板计算的结果
	tags            map[string]interface{} // 处理tag
}

// SetCount 设置模板的计算count
func (t *Template) SetCount(count int64) {
	t.count = count
}

// GetCount 获取模板的count
func (t *Template) GetCount() int64 {
	return t.count
}

// SetResult 设置模板的结果
func (t *Template) SetResult(result *common.Result) {
	t.result = result
}

func (t *Template) HasResult() bool {
	if t.result == nil {
		return false
	} else {
		return true
	}

}

// GetResult 获取导出结果
func (t *Template) GetResult() *common.Result {
	return t.result
}

// GetBeforePlugins 处理执行前参数
func (t *Template) GetBeforePlugins() []Plugin {
	return t.RouterAllConfig.BeforePlugin
}

// GetAfterPlugins 处理执行后参数
func (t *Template) GetAfterPlugins() []Plugin {
	return t.RouterAllConfig.AfterPlugin
}

/**
*  ParseRouterAll 转换总路由
 */
func (t *Template) ParseRouterAll(filePath string) (RouterAll, bool) {
	currentDir := utils.ParentDirName(filePath)
	routerALL := RouterAll{}
	routerALL.Path = filePath
	routerALL.CurrentDir = currentDir

	//加载项目路由
	msg, success := t.ParseYaml(filePath, &routerALL)
	if !success {
		t.LogErr(msg)
	}
	services := routerALL.Services

	for i := 0; i < len(services); i++ { // 处理第一层，加载项目层
		projectService := services[i]
		path := currentDir + projectService.Path
		folderPath := utils.ParentDirName(path)
		pl := ProjectLink{}
		// 设置文件夹路径
		pl.CurrentDir = folderPath
		msg, success = t.ParseYaml(path, &pl)
		// 如果项目文件转换失败
		if !success {
			t.LogErr(msg)
			continue
		}
		if utils.IsValueEmpty(projectService.Key) {
			t.LogErr(projectService)
			t.LogErr("项目没有找到配置【key】")
			continue
		}
		// 添加项目
		routerALL.AddProject(projectService.Key)
		//todo 配置关联到服务
		// 解析服务配置

		for j := 0; j < len(pl.Service); j++ { //处理第二层 加载项目里面的文件夹层
			pls := pl.Service[j]
			serviceIndexPath := folderPath + pls.Path
			serviceDir := utils.ParentDirName(serviceIndexPath)
			serviceList := ServiceList{}
			// 设置文件路径
			serviceList.Path = serviceIndexPath
			// 设置文件夹路径
			serviceList.CurrentDir = serviceDir
			msg, success = t.ParseYaml(serviceIndexPath, &serviceList)
			// 如果服务转换失败
			if !success {
				t.LogErr(msg)
				continue
			}
			for k := 0; k < len(serviceList.Service); k++ { // 处理第三层 加载文件里面的服务
				service := serviceList.Service[k]
				// 设置服务路径
				service.Path = serviceIndexPath
				// 设置服务目录
				service.CurrentDir = serviceDir
				// 设置项目
				service.Project = projectService.Key
				// 设置服务全称
				service.Service = projectService.Key + "." + service.Key
				// 添加项目服务
				routerALL.AddProjectService(projectService.Key, &service)
			}

		}

	}
	return routerALL, true
}

/*
* 转换yaml 文件
*

 */
func (t *Template) ParseYaml(filePath string, out interface{}) (interface{}, bool) {
	data, success := t.ReadFileBytes(filePath)
	if !success {
		return data, success
	}

	err := yaml.Unmarshal(data, out)
	if err != nil {
		log.Println("【" + filePath + "】转换错误")
		log.Fatalf("error: %v", err)
		return err.Error(), false
	}
	return out, true

}

/**
** 获取应用的key
 */
func (t *Template) ReadFileBytes(filePath string) ([]byte, bool) {

	content, success := utils.ReadFileBytes(filePath)
	return content, success
}

/**
** 获取应用的key
 */
func (t *Template) GetAppKey(key string) string {

	appKeyValue := utils.GetAppKey(key)
	if utils.IsValueEmpty(appKeyValue) {
		t.LogErr("没有找到配置" + key)
	}
	return appKeyValue

}

func (t *Template) LogData(data interface{}) {
	// :todo 写日志
	log.Printf("%#v", t.getLogData(data))
}
func (t *Template) LogErr(data interface{}) {
	// :todo 写日志

	log.Printf("%#v", t.getLogData(data))
}
func (t *Template) getLogData(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	return utils.GetJSONData(data)
}

//获取服务名称
func (t *Template) GetService() string {

	name := t.Service
	if !utils.IsValueEmpty(name) {
		return name
	}
	n := t.paramPool[ServiceName]
	return utils.Strval(n)

}

func (t *Template) SetService(serviceName string) Template {
	t.Service = serviceName
	return *t
}

//获取操作用户
func (t *Template) GetOpUser() string {
	return t.OpUser
}

func (t *Template) SetOpUser(opUser string) Template {
	t.OpUser = opUser
	return *t
}

// 获取事件ID
func (t *Template) GetEventId() string {
	return t.EventId
}

/**
* 获取结果
 */
func (t *Template) Result() {

}
func (t *Template) SetEventId(eventId string) Template {
	t.EventId = eventId
	return *t
}

// 设置文件内容
func (t *Template) SetFileData(fileData string) Template {

	t.FileData = fileData
	return *t

}

// 获取文件内容
func (t *Template) GetFileData() string {
	return t.FileData
}
func (t *Template) AddParam(name string, data interface{}) {
	t.paramPool[name] = data
}
func (t *Template) HasParam(key string) bool {
	params := t.GetParams()
	_, ok := params[key]
	return ok
}

func (t *Template) GetParam(key string) interface{} {
	params := t.GetParams()
	param, _ := params[key]
	return param
}
func (t *Template) SetParam(key string, value interface{}) {
	params := t.GetParams()
	params[key] = value
}

//获取参数
func (t *Template) GetParams() map[string]interface{} {
	return t.paramPool
}

// 设置参数
func (t *Template) SetParams(params map[string]interface{}) Template {
	t.paramPool = params
	return *t
}

func (t *Template) ToString() string {
	b, _ := json.Marshal(t)
	return string(b)
}
