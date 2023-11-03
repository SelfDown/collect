package collect

import (
	utils "collect.mod/src/collect/utils"
	"encoding/json"
)

/*
* 加载data_file 文件内容，处理require 文件的引用
* 将服务的文件路径转换成文件内容
 */
func (t *PluginLoader) LoadModifyConfig(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		if !utils.IsValueEmpty(service.ModifyConfig) {
			var config ModifyConfig
			json.Unmarshal([]byte(service.ModifyConfigContent), &config)
			service.ModifyConfigData = &config

		}
	}

}
