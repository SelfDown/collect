package collect

// import (
// 	"reflect"

// 	service_config "test.mod/src/collect/service_config"
// 	template "test.mod/src/collect/template"
// 	utils "test.mod/src/collect/utils"
// 	pluginUtils "test.mod/src/plugin/plugin_utils"
// 	startupPlugin "test.mod/src/plugin/startup_plugin"
// )

// /*
// * 项目总路由
//  */
// var localRouter service_config.RouterAll

// const SERVICE_NAME = "service"

// /**
// * 根据参数转模板服务
//  */
// func NewTemplateService(params map[string]interface{}) Result {
// 	r := Result{}
// 	service := utils.Strval(params[SERVICE_NAME])
// 	if utils.IsValueEmpty(service) {
// 		return *r.NotOk("服务没有找到")
// 	}
// 	scfg, success, errMsg := localRouter.GetProjectService(service)
// 	if !success {
// 		return *r.NotOk(errMsg)
// 	}

// 	return *r.Ok(scfg, "查询成功")

// }

// func LoadSystemServices(t *template.Template) {
// 	if !reflect.DeepEqual(localRouter, service_config.RouterAll{}) {
// 		return
// 	}
// 	collect_file_path := t.GetAppKey("collect_file_path")
// 	routerALL, success := t.ParseRouterAll(collect_file_path)
// 	if !success {
// 		return
// 	}
// 	var loader startupPlugin.PluginLoader
// 	pluginUtils.LoadTemplatePlugins(&loader, routerALL.LoadStartupPlugin, *t, &routerALL)
// 	localRouter = routerALL
// }
