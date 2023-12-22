package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	"reflect"
)

/*
* 加载模板级别的插件
 */
func LoadTemplatePlugins(pluginLoader interface{}, plugins []Plugin, t *Template, routerAll *RouterAll) {
	for _, plugin := range plugins {
		_callPluginFunc(pluginLoader, plugin, t, routerAll)
	}
}
func CallPluginFunc(pluginLoader interface{}, plugin Plugin, t *Template, args ...interface{}) *common.Result {
	result := _callPluginFunc(pluginLoader, plugin, t, t.RouterAllConfig, args...)
	return result
}

/**
* 调用启动插件方法
* @pluginLoader 插件接收体
* @plugin       插件对象
* @t            模板
* @routerAll    总路由
 */

func _callPluginFunc(pluginLoader interface{}, plugin Plugin, t *Template, routerAll *RouterAll, args ...interface{}) *common.Result {

	rf := reflect.ValueOf(pluginLoader)
	rft := rf.Type()
	parameter := make([]reflect.Value, 0)
	fname := plugin.Method
	parameter = append(parameter, reflect.ValueOf(plugin))
	parameter = append(parameter, reflect.ValueOf(t))
	parameter = append(parameter, reflect.ValueOf(routerAll))
	for _, arg := range args {
		parameter = append(parameter, reflect.ValueOf(arg))
	}
	_, success := rft.MethodByName(fname)
	if !success {
		msg := "【" + fname + "】方法不存在！！！"
		t.LogErr(msg)
		return common.NotOk(msg)
	}
	result := rf.MethodByName(fname).Call(parameter)
	if result == nil {
		return common.Ok(nil, "成功")
	}
	v := result[0]
	d := v.Interface().(*common.Result)
	return d

}
