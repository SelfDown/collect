package collect

import (
	"collect.mod/model"
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"database/sql"
	"github.com/demdxx/gocast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"reflect"
	"runtime"
	"strings"
)

var db0 *sql.DB
var gormDb *gorm.DB

type BaseHandler struct {
}

func (s *BaseHandler) RunFuncName() string {

	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func (s *BaseHandler) GetGormDb() *gorm.DB {
	//如果连接过则直接，返回对象
	if gormDb != nil {
		return gormDb
	}
	db, _ := s.GetDatasource()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		CreateBatchSize: 1000,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	return gormDB
}

// GetDatasource 获取数据库连接
func (s *BaseHandler) GetDatasource() (*sql.DB, error) {
	if db0 != nil {
		return db0, nil
	}
	// 获取连接信息
	dataSourceName := utils.GetAppKey("dataSourceName")
	// 获取驱动
	driverName := utils.GetAppKey("driverName")

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return nil, err
	}
	db.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return nil, err
	}
	db0 = db
	return db, err
}

func (s *BaseHandler) Result(template *config.Template, ts *TemplateService) *common.Result {
	return common.Ok(nil, "")
}

func (s *BaseHandler) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	return common.Ok(nil, "")
}
func (s *BaseHandler) UpdateFields(params map[string]interface{}, modelData interface{}, ignoreFields []string, updateFields []string) (interface{}, []string) {
	_, names := utils.SetDataValueByParams(params, modelData, ignoreFields, updateFields)
	return modelData, names
}
func (s *BaseHandler) UpdateFieldsToMap(params map[string]interface{}, modelData interface{}, ignoreFields []string, updateFields []string) (map[string]interface{}, []string) {
	data, names := utils.SetDataValueByParams(params, modelData, ignoreFields, updateFields)
	return data, names
}
func (s *BaseHandler) UpdateFieldsToMapList(models []map[string]interface{}, modelData interface{}, template *config.Template) ([]map[string]interface{}, []string) {
	modelList := make([]map[string]interface{}, 0)
	var fieldNames []string
	for _, item := range models {
		modelItem := model.CloneModel(template.Table)
		dataItem, names := s.UpdateFieldsToMap(item, &modelItem, template.IgnoreFields, template.UpdateFields)
		modelList = append(modelList, dataItem)
		if fieldNames == nil {
			fieldNames = names
		}
	}
	return modelList, fieldNames
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
func (s *BaseHandler) HandlerFilter(template *config.Template) (interface{}, []interface{}) {
	query := ""
	whereList := make([]string, 0)
	valueList := make([]interface{}, 0)
	params := template.GetParams()
	for key, paramKey := range template.Filter {
		value := utils.RenderVarOrValue(paramKey, params)
		// 如果参数中没有定义此值则，取配置的值
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
func (s *BaseHandler) CheckFilter(template *config.Template, model interface{}) *common.Result {
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
