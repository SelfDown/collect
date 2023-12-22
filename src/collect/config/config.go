package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

/*
* 项目总路由
 */
var localRouter RouterAll

const ServiceName = "service"

//func GetModuleRegister(name string) module_result.ModuleResult {
//	return localRouter.GetModuleRegister(name)
//}

/*
* 获取服务名称
 */
func GetServiceName(params map[string]interface{}) string {
	service := utils.Strval(params[ServiceName])
	return service

}

/**
* 根据参数转模板服务
 */
func NewTemplateService(params map[string]interface{}) *common.Result {
	r := common.Result{}
	service := utils.Strval(params[ServiceName])
	if utils.IsValueEmpty(service) {
		return r.NotOk("服务没有找到")
	}
	scfg, success, errMsg := localRouter.GetProjectService(service)

	if !success {
		return r.NotOk(errMsg)
	}

	return r.Ok(scfg, "查询成功")

}
func SetLocalRouter(routerAll RouterAll) {
	localRouter = routerAll
}

func GetLocalRouter() *RouterAll {
	return &localRouter
}

// func SetRouterRegisterList(registerList []ModuleResult) {
// 	localRouter.SetRegisterList(registerList)
// }

//func GetRouterRegisterList() []module_result.ModuleResult {
//	return localRouter.GetRegisterList()
//}

// func LoadSystemServices(t *Template) {
// 	if !reflect.DeepEqual(localRouter, RouterAll{}) {
// 		return
// 	}
// 	collect_file_path := t.GetAppKey("collect_file_path")
// 	routerALL, success := t.ParseRouterAll(collect_file_path)
// 	if !success {
// 		return
// 	}
// 	var loader PluginLoader
// 	LoadTemplatePlugins(&loader, routerALL.LoadStartupPlugin, *t, &routerALL)
// 	localRouter = routerALL
// }
