package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	serviceImp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
	"github.com/demdxx/gocast"
	"reflect"
	"strings"
)

type ModelUpdateService struct {
	serviceImp.BaseHandler
}

func getOp(op string, value interface{}) (string, interface{}) {
	if utils.IsValueEmpty(op) { // 处理等于
		return "=", value
	}
	op = strings.ToLower(op)
	if op == "in" {
		return op, value
	} else if op == "isnull" {
		b := gocast.ToBool(value)
		if b {
			return "is null ", nil
		} else {
			return "is not null", nil
		}

	}
	return "=", value
}
func GetFieldName(key string) string {
	fieldOp := strings.Split(key, "__")
	return fieldOp[0]
}
func GetOpName(key string) string {
	fieldOp := strings.Split(key, "__")
	if len(fieldOp) >= 2 {
		return fieldOp[1]
	}
	return ""
}
func HandlerFilter(template *config.Template) (interface{}, []interface{}) {
	query := ""
	whereList := make([]string, 0)
	valueList := make([]interface{}, 0)
	for key, paramKey := range template.Filter {
		var value interface{}
		if reflect.TypeOf(paramKey).String() == "string" {
			value = template.GetParam(utils.Strval(paramKey))
		}
		// 如果参数中没有定义此值则，取配置的值
		if value == nil {
			value = paramKey
		}
		field := GetFieldName(key)
		op := GetOpName(key)
		opNew, valueNew := getOp(op, value)
		temp := field + " " + opNew
		if valueNew != nil {
			temp += " ? "
		}
		whereList = append(whereList, temp)
		if valueNew == nil { // 不允许加null
			continue
		}
		valueList = append(valueList, valueNew)

	}
	query = strings.Join(whereList, " AND ")
	return query, valueList
}
func CheckFilter(template *config.Template, model interface{}) *common.Result {
	if template.Filter == nil {
		return common.NotOk("过滤条件不能为空")
	}
	modelValue := reflect.ValueOf(model)
	for key, paramKey := range template.Filter {
		if utils.IsValueEmpty(key) {
			return common.NotOk("存在过滤key为空")
		}
		if utils.IsValueEmpty(paramKey) {
			return common.NotOk("存在过滤[" + key + "]值value为空")
		}
		//检查key中是否包含空格
		if strings.Contains(key, " ") {
			return common.NotOk("存在过滤[" + key + "]包含空格")
		}
		//检查key中是否包含;防止sql注入
		if strings.Contains(key, ";") {
			return common.NotOk("存在过滤[" + key + "]包含[;]")
		}
		fieldOrgName := GetFieldName(key)
		fieldName := utils.ToSchemaName(fieldOrgName)
		field := modelValue.FieldByName(fieldName)
		if !field.IsValid() {
			return common.NotOk("存在过滤[" + key + "]中数据库字段[" + fieldOrgName + "]不存在")
		}

	}
	return common.Ok(nil, "检查成功")
}
func (s *ModelUpdateService) Result(template *config.Template) *common.Result {

	params := template.GetParams()
	tableName := template.Table
	modelData := model.GetModel(tableName)
	checkResult := CheckFilter(template, modelData)
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
	pk := model.GetPrimaryKey(tableName)
	//不允许更新主键
	for _, k := range pk {
		if !utils.StringArrayContain(template.IgnoreFields, k) {
			template.IgnoreFields = append(template.IgnoreFields, k)
		}
	}
	_, fieldNames := s.UpdateFields(params, &modelData, template.IgnoreFields, template.UpdateFields)
	gormDB := s.GetGormDb()
	//生成where 条件+参数
	query, args := HandlerFilter(template)
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
