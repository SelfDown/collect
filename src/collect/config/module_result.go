package collect

import (
	common "collect.mod/src/collect/common"
	// template "test.mod/src/collect/template"
)

type ModuleResult interface {
	Result(template *Template) (*common.Result, error)
}
