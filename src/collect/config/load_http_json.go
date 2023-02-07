package collect

import (
	utils "collect.mod/src/collect/utils"
	"encoding/json"
	text_template "text/template"
)

/*
* 加载http_json 文件内容，处理require 文件的引用
* 将服务的文件路径转换成文件内容
 */
func (t *PluginLoader) LoadHttpJson(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		if !utils.IsValueEmpty(service.HttpJson) {
			var config HttpConfig
			json.Unmarshal([]byte(service.HttpJsonContent), &config)
			// 处理http url 地址
			if !utils.IsValueEmpty(config.Url) {
				tpl, err := _load_template(config.Url)
				if err != nil {
					template.LogData(err)
					continue
				}
				config.UrlTpl = tpl
			}
			// 处理head
			if !utils.IsValueEmpty(config.Header) {
				config.HeaderTpl = make(map[string]*text_template.Template)
				for k, v := range config.Header {
					tpl, err := _load_template(v)
					if err != nil {
						template.LogData(err)
						continue
					}
					config.HeaderTpl[k] = tpl
				}
			}
			if !utils.IsValueEmpty(config.Data) {
				data := config.Data
				str, err := json.Marshal(data)
				if err != nil {
					template.LogData(err)
					continue
				}
				tpl, err := _load_template(string(str))
				if err != nil {
					template.LogData(err)
					continue
				}
				config.DataTpl = tpl
			}
			//todo 处理data,处理json模板渲染

			service.HttpConfigData = &config

		}
	}

}
