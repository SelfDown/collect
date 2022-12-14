package collect

import (
	// "fmt"

	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	// sql "test.mod/src/collect/service_imp/sql"
)

func LoadSystemServices() config.RouterAll {
	t := config.Template{}
	utils.LoadAppProperties("./conf/application.properties")
	collectFilePath := t.GetAppKey("collect_file_path")
	routerALL, success := t.ParseRouterAll(collectFilePath)
	if !success {
		return config.RouterAll{}
	}
	var loader config.PluginLoader
	//加载启动插件
	config.LoadTemplatePlugins(&loader, routerALL.LoadStartupPlugin, &t, &routerALL)
	return routerALL
}
