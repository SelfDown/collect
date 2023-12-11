package collect

import (
	"collect/model"
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type ModelDeleteService struct {
	BaseHandler
}

func (s *ModelDeleteService) Result(template *config.Template, ts *TemplateService) *common.Result {

	//params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	checkResult := s.CheckFilter(template, modelData)
	if !checkResult.Success {
		return checkResult
	}
	if modelData == nil {
		return common.NotOk(tableName + "没有找到，请检查模型数据")
	}

	gormDB := s.GetGormDb()
	//生成where 条件+参数
	query, args := s.HandlerFilter(template)
	//执行
	if template.Log {
		template.LogData("删除表[" + tableName + "]")
		template.LogData("过滤条件[" + utils.Strval(query) + "]")
		template.LogData("过滤参数:")
		template.LogData(args)
	}
	dbx := gormDB.Where(query, args...).Delete(modelData)
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
	return common.OkWithCount(r, "删除成功", affected)
}
