package collect

import (
	config "collect.mod/src/collect/config"
	handlerParams "collect.mod/src/collect/service_imp/handler_params"
	bulkCreate "collect.mod/src/collect/service_imp/module/bulk_create"
	modelUpsert "collect.mod/src/collect/service_imp/module/bulk_upsert"
	modelDelete "collect.mod/src/collect/service_imp/module/model_delete"
	modelSave "collect.mod/src/collect/service_imp/module/model_save"
	modelUpdate "collect.mod/src/collect/service_imp/module/model_update"
	sql "collect.mod/src/collect/service_imp/module/sql"
)

/**
* 获取注册列表，对象名称绑定key，一定要注意对象名称一定要唯一
**/
func GetRegisterList() []config.ModuleResult {
	l := make([]config.ModuleResult, 0)
	l = append(l, &sql.SqlService{})
	l = append(l, &handlerParams.UpdateField{})
	l = append(l, &handlerParams.UpdateArray{})
	l = append(l, &modelSave.ModelSaveService{})
	l = append(l, &modelUpdate.ModelUpdateService{})
	l = append(l, &modelDelete.ModelDeleteService{})
	l = append(l, &bulkCreate.BulkCreateService{})
	l = append(l, &modelUpsert.BulkUpsertService{})
	return l
}
