package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"gorm.io/gorm/clause"
)

type BulkUpsertService struct {
	BaseHandler
}

func (s *BulkUpsertService) Result(template *config.Template, ts *TemplateService) *common.Result {

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
	if template.UpdateFields != nil { // 添加主键
		primaryKeys := model.GetPrimaryKey(template.Table)
		for _, pk := range primaryKeys {
			if !utils.StringArrayContain(template.UpdateFields, pk) {
				template.UpdateFields = append(template.UpdateFields, pk)
			}
		}
	}
	//执行
	modelList, fieldNames := s.UpdateFieldsToMapList(dataList, modelData, template)
	//保存
	gormDB := s.GetGormDb()
	if template.Log {
		template.LogData("更新或新增表[" + tableName + "]")
		template.LogData("列表数据")
		template.LogData(modelList)
		template.LogData("更新字段")
		template.LogData(fieldNames)
	}
	dbx := gormDB.Model(modelData).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns(fieldNames),
	}).Create(modelList)
	affected := dbx.RowsAffected
	err := dbx.Error
	if err != nil {
		msg := err.Error()
		template.LogData(msg)
		return common.NotOk(msg)
	}
	return common.OkWithCount(params, "批量新增或修改成功", affected)
}
