package collect

import (
	utils "github.com/SelfDown/collect/src/collect/utils"
)

/*
* 加载data_file 文件内容，处理require 文件的引用
* 将服务的文件路径转换成文件内容
 */
func (t *PluginLoader) LoadProxyConfig(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for index, service := range serviceList {
		if utils.IsValueEmpty(service.Proxy.Enable) {
			continue
		}
		proxy := service.Proxy
		if !utils.IsValueEmpty(proxy.Enable) {
			tpl, err := _load_template(proxy.Enable)
			serviceList[index].Proxy.EnableTpl = tpl
			if err != nil {
				template.LogData(err)
				continue
			}
		}

	}

}
