package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type BulkCreateService struct {
	BaseHandler
}

func (s *BulkCreateService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := ts.GetModel(tableName)
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	modelField := template.ModelField
	if utils.IsValueEmpty(modelField) {
		return common.NotOk(template.GetService() + "没有配置model_field")
	}
	dataList, errMsg := utils.RenderVarToArrMap(modelField, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	if len(dataList) <= 0 {
		return common.NotOk(template.GetService() + "列表" + modelField + "数据为空")
	}
	// 将参数列表，转成模型列表
	modelList, _, errMsg := s.UpdateFieldsToMapList(dataList, modelData, template)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
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
