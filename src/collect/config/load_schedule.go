package collect

import (
	utils "collect/src/collect/utils"
)

/*
* 加载定时任务是否启动
 */
func (t *PluginLoader) LoadSchedule(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		if !utils.IsValueEmpty(service.Schedule.Enable) {
			tplName, err := _load_template(service.Schedule.Enable)
			if err != nil {
				template.LogData(err)
				continue
			}
			service.Schedule.EnableTpl = tplName

		}
	}

}
