package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	serviceImp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
)

type BulkCreateService struct {
	serviceImp.BaseHandler
}

func (s *BulkCreateService) Result(template *config.Template) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	modelField := template.ModelField
	if utils.IsValueEmpty(modelField) {
		return common.NotOk(template.GetService() + "没有配置model_field")
	}
	models := params[modelField]
	dataList := models.([]map[string]interface{})
	if len(dataList) <= 0 {
		return common.NotOk(template.GetService() + "列表" + modelField + "数据为空")
	}
	// 将参数列表，转成模型列表
	modelList, _ := s.UpdateFieldsToMapList(dataList, modelData, template)
	//保存
	gormDB := s.GetGormDb()
	dbx := gormDB.Model(modelData).Create(modelList)
	affected := dbx.RowsAffected
	err := dbx.Error
	if err != nil {
		msg := err.Error()
		template.LogData(msg)
		return common.NotOk(msg)
	}
	return common.OkWithCount(params, "批量新增成功", affected)
}
