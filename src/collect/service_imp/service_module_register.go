package collect

import (
	config "github.com/SelfDown/collect/src/collect/config"
	"reflect"
)

// 注册的插件
var localModuleDict map[string]config.Plugin

// 注册的模块
var localModuleRegisterDict map[string]ModuleResult
var outerRegisterModule []ModuleResult

// SetOuterModuleRegister 设置外部处理器
func SetOuterModuleRegister(registerList []ModuleResult) {
	moduleDict := localModuleDict
	outerRegisterModule = registerList
	for _, reg := range registerList {
		// 这里根据字符串，注册层服务
		name := reflect.TypeOf(reg).Elem().Name()
		//moduleRegisterDict 进行赋值
		if module, ok := moduleDict[name]; ok {
			localModuleRegisterDict[module.Key] = reg
		}

	}
}

func SetRegisterList(t *config.RouterAll, registerList []ModuleResult) {

	//	// 设置转换字典
	moduleDict := make(map[string]config.Plugin)
	for _, module := range t.ModuleHandler {
		moduleDict[module.Path] = module
	}
	for _, module := range t.DataHandler {
		moduleDict[module.Path] = module
	}

	localModuleDict = moduleDict
	// 清空字典
	localModuleRegisterDict = make(map[string]ModuleResult)
	for _, reg := range registerList {
		// 这里根据字符串，注册层服务
		name := reflect.TypeOf(reg).Elem().Name()
		//moduleRegisterDict 进行赋值
		if module, ok := moduleDict[name]; ok {
			localModuleRegisterDict[module.Key] = reg
		}

	}

}

func GetModuleRegister(name string) ModuleResult {
	//	return localRouter.GetModuleRegister(name)
	return localModuleRegisterDict[name]
}
