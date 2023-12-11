package collect

import (
	"strings"

	"regexp"

	utils "collect/src/collect/utils"
)

/**
* 递归处理文件依赖
**/
func _handler_require_file(reg string, fileContent string, dir string) (string, bool) {
	re := regexp.MustCompile(reg)
	found := re.FindAllStringSubmatch(fileContent, -1)
	for _, require_file := range found {
		// 原始占位符
		orignial := require_file[0]
		// 替换掉' 和" 以获取文件名
		require_file_name := strings.ReplaceAll(require_file[1], "'", "")
		require_file_name = strings.ReplaceAll(require_file_name, "\"", "")
		//读取文件内容
		requirefileContent, success := utils.ReadFileContent(dir + require_file_name)
		// 如果读取到文件，直接替换原始占位符
		if success {
			fileContent = strings.ReplaceAll(fileContent, orignial, requirefileContent)
		} else {
			return requirefileContent, success
		}
		// 判断是否多重引用
		found_child := re.FindAllStringSubmatch(fileContent, -1)
		// 递归处理子文件依赖
		if len(found_child) > 0 {
			fileContent, success = _handler_require_file(reg, fileContent, dir)
			if !success {
				return fileContent, success
			}
		}

	}
	return fileContent, true
}

/**
* 加载插件文件，处理require
 */
func _handler_plugins_file(dataFile string, dir string, plugins []Plugin) (string, bool) {
	fileContent, success := utils.ReadFileContent(dir + dataFile)
	for _, plugin := range plugins {

		//todo  不同的插件，调用不同的方法，目前只有一个require ，后续也不知道是否有扩展，先写死方法
		fileContent, success = _handler_require_file(plugin.Reg, fileContent, dir)
		if !success {
			return fileContent, success
		}

	}
	return fileContent, success

}

/*
* 加载data_file 文件内容，处理require 文件的引用
* 将服务的文件路径转换成文件内容
 */
func (t *PluginLoader) LoadDataFile(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		for _, field := range config.Fields { // 将里面的字段转换成模板
			// 根据from 获取来源字段名称
			dataFrom := utils.GetDataValueStr(field.From, *service)
			if utils.IsValueEmpty(dataFrom) {
				continue
			}
			// 根据来源文件转，文件内容
			fileContent, success := _handler_plugins_file(dataFrom, service.CurrentDir, routerAll.FileContentPlugin)
			if !success { // 如果失败，则打印日志
				template.LogData(fileContent)
				continue
			}
			// 设置文件内容
			utils.SetDataValue(field.To, fileContent, service)
		}

	}

}
