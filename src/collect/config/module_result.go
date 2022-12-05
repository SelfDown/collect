package collect

import (
	common "collect.mod/src/collect/common"
	// template "test.mod/src/collect/template"
)

type ModuleResult interface {
	//执行结果
	Result(template *Template) *common.Result
	// 处理数据
	HandlerData(template *Template, handlerParam *HandlerParam) *common.Result
}
