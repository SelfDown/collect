package collect

import (
	"collect/model"
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type ModelSaveService struct {
	BaseHandler
}

func (s *ModelSaveService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	fieldOptions, errMsg := s.getFieldOptions(template.Options, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	s.UpdateFields(params, &modelData, template.IgnoreFields, template.UpdateFields, fieldOptions)
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
