package collect

import (
	"reflect"
	"strconv"
	"strings"
)

type NormalArray struct {
	NormalParam
}

const ArrSplit = "_______________COLLECT_SQL_ARR_PARAM_SPLIT________________"

func (t *NormalArray) oriArrKey(index int) string {
	return t.ParamKey + ArrSplit + strconv.Itoa(index) + ArrSplit
}

func (t *NormalArray) get_arr_key(index int) string {
	return "{{." + t.oriArrKey(index) + "}}"
}

// 获取数组的参数变量
func (t *NormalArray) GetSqlParamKeyValue(to_param_key bool) interface{} {
	var param_key_value string
	if to_param_key {
		l := make([]string, 0)
		for i := 0; i < reflect.ValueOf(t.ParamValue).Len(); i++ {
			arr_key := t.get_arr_key(i)
			l = append(l, arr_key)
		}
		return strings.Join(l, ",")

	}
	return param_key_value
}

// /**
// * 获取数组的实际值
// **/
// func (t *NormalArray) GetValue() []interface{} {

// 	value := t.ParamValue
// 	if value == nil {
// 		return make([]interface{}, 0)
// 	}
// 	// 反射数组的类型
// 	sv := reflect.ValueOf(value)
// 	size := sv.Len()
// 	result := make([]interface{}, size)
// 	for i := 0; i < size; i++ {
// 		// 获取实际值
// 		result = append(result, sv.Index(i).Interface())
// 	}
// 	return result
// }

/**
* 获取数组对象的key
 */
func (t *NormalArray) GetSqlParamParamKeyList() []AttrParam {

	apList := make([]AttrParam, 0)
	for i := 0; i < reflect.ValueOf(t.ParamValue).Len(); i++ {
		arr_key := t.oriArrKey(i)
		// 将数组中每个值，转换成独立变量
		ap := AttrParam{AttrName: arr_key, AttrValue: reflect.ValueOf(t.ParamValue).Index(i).Interface()}
		apList = append(apList, ap)
	}
	return apList
}
