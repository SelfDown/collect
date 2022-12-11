package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	serviceImp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
)

type ModelSaveService struct {
	serviceImp.BaseHandler
}

func updateFields(params map[string]interface{}, modelData interface{}, ignoreFields []string, updateFields []string) interface{} {
	utils.SetDataValueByParams(params, modelData, ignoreFields, updateFields)
	return modelData
}

func (s *ModelSaveService) Result(template *config.Template) *common.Result {
	params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	updateFields(params, &modelData, template.IgnoreFields, template.UpdateFields)
	gormDB := s.GetGormDb()
	gormDB.Create(modelData)
	return common.Ok(params, "保存成功")
}
