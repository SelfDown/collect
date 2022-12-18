package collect

import (
	"encoding/json"
	"log"

	// service_config "test.mod/src/collect/service_config"
	utils "collect.mod/src/collect/utils"

	"gopkg.in/yaml.v2"
)

type Template struct {
	Service   string                 // 服务名称
	OpUser    string                 // 操作用户
	EventId   string                 // 事件ID
	paramPool map[string]interface{} //请求参数池
	ServiceConfig
	RouterAllConfig *RouterAll // 总服务路由
	
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
* 转换总路由
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
		folder_path := utils.ParentDirName(path)
		pl := ProjectLink{}
		// 设置文件夹路径
		pl.CurrentDir = folder_path
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
			service_index_path := folder_path + pls.Path
			service_dir := utils.ParentDirName(service_index_path)
			serviceList := ServiceList{}
			// 设置文件路径
			serviceList.Path = service_index_path
			// 设置文件夹路径
			serviceList.CurrentDir = service_dir
			msg, success = t.ParseYaml(service_index_path, &serviceList)
			// 如果服务转换失败
			if !success {
				t.LogErr(msg)
				continue
			}
			for k := 0; k < len(serviceList.Service); k++ { // 处理第三层 加载文件里面的服务
				service := serviceList.Service[k]
				// 设置服务路径
				service.Path = service_index_path
				// 设置服务目录
				service.CurrentDir = service_dir
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
* todo 这里需要将int 转int32，比如statu=1，类型是int，但是数据库生成是int32，保存失败

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
	log.Println(t.getLogData(data))
}
func (t *Template) LogErr(data interface{}) {
	// :todo 写日志

	log.Println(t.getLogData(data))
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
	n := t.paramPool[SERVICE_NAME]
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
