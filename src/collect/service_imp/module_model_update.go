package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type ModelUpdateService struct {
	BaseHandler
}

func (s *ModelUpdateService) Result(template *config.Template, ts *TemplateService) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := ts.GetModel(tableName)
	checkResult := s.CheckFilter(template, modelData)
	if !checkResult.Success {
		return checkResult
	}
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}
	// 修改数据
	if template.IgnoreFields == nil {
		template.IgnoreFields = make([]string, 0)
	}
	pk := ts.GetPrimaryKey(tableName)
	//不允许更新主键
	for _, k := range pk {
		if !utils.StringArrayContain(template.IgnoreFields, k) {
			template.IgnoreFields = append(template.IgnoreFields, k)
		}
	}
	fieldOptions, errMsg := s.getFieldOptions(template.Options, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	_, fieldNames := s.UpdateFields(params, &modelData, template.IgnoreFields, template.UpdateFields, fieldOptions)
	gormDB := s.GetGormDb()
	//生成where 条件+参数
	query, args := s.HandlerFilter(template)
	//执行
	if template.Log {
		template.LogData("更新表[" + tableName + "]")
		template.LogData("过滤条件[" + utils.Strval(query) + "]")
		template.LogData("过滤参数:")
		template.LogData(args)
		template.LogData("传入数据:")
		template.LogData(modelData)
		template.LogData("更新字段:")
		template.LogData(fieldNames)
	}
	dbx := gormDB.Where(query, args...).Select(fieldNames).Updates(modelData)
	affected := dbx.RowsAffected
	err := dbx.Error
	if err != nil {
		return common.NotOk(err.Error())
	}
	r := make(map[string]interface{})
	r["affected"] = affected
	if template.Log {
		template.LogData("影响行数:" + utils.Strval(affected))
	}
	return common.OkWithCount(r, "修改成功", affected)
}
