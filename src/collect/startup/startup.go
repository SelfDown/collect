package collect

import (
	// "fmt"

	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
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
	// 初始化数据连接
	//base := collect.BaseHandler{}
	//base.GetDatasource()
	return routerALL
}
