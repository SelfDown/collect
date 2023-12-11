package collect

import (
	utils "collect/src/collect/utils"
	"encoding/json"
)

/*
* 加载data_json 文件内容，
 */
func (t *PluginLoader) LoadDataJson(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		if !utils.IsValueEmpty(service.DataJson) {
			var config DataJsonConfig
			json.Unmarshal([]byte(service.DataJsonContent), &config)
			service.DataJsonConfig = &config
		}
	}

}
