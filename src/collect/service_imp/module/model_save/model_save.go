package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	serviceImp "collect.mod/src/collect/service_imp"
)

type ModelSaveService struct {
	serviceImp.BaseHandler
}

func (s *ModelSaveService) Result(template *config.Template) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	s.UpdateFields(params, &modelData, template.IgnoreFields, template.UpdateFields)
	gormDB := s.GetGormDb()
	dbx := gormDB.Create(modelData)
	affected := dbx.RowsAffected
	err := dbx.Error
	if err != nil {
		msg := err.Error()
		template.LogData(msg)
		return common.NotOk(msg)
	}
	return common.OkWithCount(params, "保存成功", affected)
}
