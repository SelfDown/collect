package collect

import (
	"encoding/json"
	utils "github.com/SelfDown/collect/src/collect/utils"
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

			for index, item := range config.Fields {
				// 转换enable
				if !utils.IsValueEmpty(item.Enable) {
					tpl, err := _load_template(item.Enable)
					if err != nil {
						template.LogData(err)
						continue
					}
					config.Fields[index].EnableTpl = tpl
				}
				if !utils.IsValueEmpty(item.IfTemplate) {
					tpl, err := _load_template(item.IfTemplate)
					if err != nil {
						template.LogData(err)
						continue
					}
					config.Fields[index].IfTemplateTpl = tpl
				}

			}
			service.ModifyConfigData = &config

		}
	}

}
