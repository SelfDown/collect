package collect

import (
	config "collect.mod/src/collect/config"
	sql "collect.mod/src/collect/service_imp/sql"
)

/**
* 获取注册列表
**/
func GetRegisterList() []config.ModuleResult {
	l := make([]config.ModuleResult, 0)
	l = append(l, &sql.SqlService{})
	return l
}
