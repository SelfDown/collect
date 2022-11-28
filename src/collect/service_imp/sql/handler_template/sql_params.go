package collect

import (
	"reflect"
	"strings"

	utils "collect.mod/src/collect/utils"
)

//todo  定义一个结构体
// 1 根据字段名返回参数情况下，字段名，值情况下的字段名。用于渲染渲染第一sql 渲染变量处理，第二次渲染预编译变量
// 2 返回字段名形式下的，参数值列表，用于第二次的渲染参数值
// 3 返回 2渲染实际结果值

/**
* to_param_key 渲染第一次模板变量为true
**/
func GetSqlParamKeyValue(param_key string, param_value interface{}, to_param_key bool) string {
	var param_key_value string
	if to_param_key {
		return param_key
	}
	if utils.IsArray(param_value) {
		size := reflect.ValueOf(param_value).Len()
		param_key_value = utils.MultiplyJoinComma("?", size)
	} else {
		param_key_value = "?"
	}
	return param_key_value
}
func GetSqlParamKeyName(param_key string) string {
	if utils.IsValueEmpty(param_key) {
		return ""
	}
	v := strings.ReplaceAll(param_key, "{{", "")
	v = strings.ReplaceAll(v, "}}", "")
	v_list := strings.Split(v, ".")
	if len(v_list) < 2 {
		return ""
	}
	return v_list[1]
}

/*
* sql 变量类型
 */
type SqlParamName struct {
	Name        string         // 变量名称
	Value       interface{}    // 值
	VarType     string         // 变量类型
	ItemName    string         // 子变量前缀
	ItemVarList []SqlParamName // 子变量列表
}

const SIMPLE_PARAM = "simple"
const LIST_SIMPLE = "list_simple"
const LIST_OBJ = "list_obj"

func newSimpleSqlParam(name string) SqlParamName {
	return SqlParamName{Name: name, VarType: SIMPLE_PARAM}
}

func newListObjSqlParam(name string, item_name string, item_var_list []SqlParamName) SqlParamName {
	return SqlParamName{Name: name, VarType: LIST_OBJ, ItemName: item_name, ItemVarList: item_var_list}
}
func (t *SqlParamName) GetName() string {
	return t.Name

}
func (t *SqlParamName) IsSimple() bool {
	return t.VarType == SIMPLE_PARAM

}

func (t *SqlParamName) GetItemName() string {
	return t.ItemName
}

func (t *SqlParamName) GetItemVarList() []SqlParamName {
	return t.ItemVarList
}
