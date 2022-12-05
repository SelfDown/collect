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
type BaseParam interface {
	// 设置字段名名
	SetParamKey(string)
	// 设置原始key
	SetParamKeyOrinal(string)
	// 设置字段值
	SetParamKeyValue(interface{})
	// 获取字段名的参数值,比如 to_param_key 为true ，直接返回字段名称，如果为false ，直接返回 sql占位符号 ?。如果是数组则是，多个? 拼接的字符串
	GetSqlParamKeyValue(to_param_key bool) interface{}
	// 获取字段名,获取模板变量里面的字段名称，比如 {{.userList}},获取变量userList
	GetSqlParamParamKey() string
	// 获取原始变量值{{.userList}},获取变量{{.userList}}
	GetParamKeyOriginal() string
	// 获取参数的值
	GetValue() []interface{}
	// 获取数组对象列表的key
	GetSqlParamParamKeyList() []AttrParam
	// 设置子变量名称
	SetChildKeys(childrenKeys []string)
}
type AttrParam struct {
	AttrName  string
	AttrValue interface{}
}

func GetSqlParamKey(sqlParamKey string) string {
	param_key := sqlParamKey
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

const PARAM_SIMPLE_PARAM = "simple"
const PARAM_LIST_SIMPLE = "list_simple"
const PARAM_LIST_OBJ = "list_obj"

/**
* 获取简单对象
 */
func GetSimpleParamObj(ParamKeyOriginal string, params map[string]interface{}) BaseParam {
	param_key := GetSqlParamKey(ParamKeyOriginal)
	if !utils.IsEmpty(param_key, params) && utils.IsArray(params[param_key]) {
		return getParamObj(ParamKeyOriginal, params, PARAM_LIST_SIMPLE)

	} else {
		return getParamObj(ParamKeyOriginal, params, PARAM_SIMPLE_PARAM)
	}

}

/**
* 获取数组对象列表
 */
func GetArrayObjParamObj(paramKey string, childrenKeys []string, params map[string]interface{}) BaseParam {
	// param_key := GetSqlParamKey(ParamKeyOriginal)
	paramArr := getParamObj(paramKey, params, PARAM_LIST_OBJ)
	// 设置字段
	paramArr.SetParamKey(paramKey)
	// 设置子变量
	paramArr.SetChildKeys(childrenKeys)
	// 设置值
	paramArr.SetParamKeyValue(params[paramKey])
	return paramArr

}

/**
* 获取模板变量
 */
func getParamObj(ParamKeyOriginal string, params map[string]interface{}, varType string) BaseParam {

	var param BaseParam = nil
	// 处理普通变量
	if varType == PARAM_SIMPLE_PARAM {
		param = &NormalParam{}
		// 处理简单数组
	} else if varType == PARAM_LIST_SIMPLE {
		param = &NormalArray{}
		// 处理数组对象
	} else if varType == PARAM_LIST_OBJ {
		param = &NormalArrayObj{}
	}
	if param == nil {
		return param
	}
	// 设置原始key
	param.SetParamKeyOrinal(ParamKeyOriginal)
	param_key := param.GetSqlParamParamKey()
	// 如果参数中存在此值，则填充值
	if !utils.IsEmpty(param_key, params) {
		param.SetParamKeyValue(params[param_key])
	}

	return param

}

func NewParamFromMap(param_key_original string, param map[string]interface{}) BaseParam {
	new_param_key := GetSqlParamKey(param_key_original)
	param_value := param[new_param_key]
	return NewParam(new_param_key, param_key_original, param_value)
}

func NewParam(param_key string, param_key_original string, param_value interface{}) BaseParam {
	var np BaseParam
	if utils.IsArray(param_value) {

		t := reflect.ValueOf(param_value)
		if t.Len() > 0 {
			tmp := make(map[string]interface{}, 0)
			if reflect.ValueOf(tmp).Kind() == t.Index(0).Kind() {
				np = &NormalArrayObj{}
			} else {
				np = &NormalArray{}
			}

		} else {
			np = &NormalArray{}
		}

	} else {
		np = &NormalParam{}
	}
	np.SetParamKey(param_key)
	np.SetParamKeyOrinal(param_key_original)
	np.SetParamKeyValue(param_value)
	return np
}
