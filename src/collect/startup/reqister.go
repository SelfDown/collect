package collect

import (
	config "collect.mod/src/collect/config"
	handler_params "collect.mod/src/collect/service_imp/handler_params"
	sql "collect.mod/src/collect/service_imp/module/sql"
)

/**
* 获取注册列表
**/
func GetRegisterList() []config.ModuleResult {
	l := make([]config.ModuleResult, 0)
	l = append(l, &sql.SqlService{})
	l = append(l, &handler_params.UpdateField{})
	return l
}
