package collect

import (
	"reflect"
	"strconv"

	utils "github.com/SelfDown/collect/src/collect/utils"
)

type NormalArrayObj struct {
	NormalParam
}

const ARR_OBJ_SPLIT = "________________COLLECT_SQL_OBJ_ARR_PARAM_SPLIT________________"

func (t *NormalArrayObj) _ori_arr_key(index int, field string) string {
	return t.ParamKey + ARR_OBJ_SPLIT + strconv.Itoa(index) + ARR_OBJ_SPLIT + field
}

func (t *NormalArrayObj) get_arr_key(index int, field string) string {

	//获取数组对象的key
	return "{{." + t._ori_arr_key(index, field) + "}}"
}

// 获取数组的参数变量
func (t *NormalArrayObj) GetSqlParamKeyValue(to_param_key bool) interface{} {

	sv := reflect.ValueOf(t.ParamValue).Interface().([]map[string]interface{})
	l := make([]map[string]interface{}, 0)
	for i, item := range sv {
		item_copy := utils.CopyMap(item)
		for _, field := range t.ChildKeys {
			item_copy[field] = t.get_arr_key(i, field)
		}
		l = append(l, item_copy)
	}

	return l
}

// /**
// * 获取数组的实际值
// **/
// func (t *NormalArrayObj) GetValue() []interface{} {

// 	value := t.ParamValue
// 	// if value == nil {
// 	// 	return make([]interface{}, 0)
// 	// }
// 	// // 反射数组的类型
// 	// sv := reflect.ValueOf(value)
// 	// size := sv.Len()
// 	// result := make([]interface{}, size)
// 	// for i := 0; i < size; i++ {
// 	// 	// 获取实际值
// 	// 	result = append(result, sv.Index(i).Interface())
// 	// }
// 	return value
// }

/**
* 获取数组对象的key
 */
func (t *NormalArrayObj) GetSqlParamParamKeyList() []AttrParam {

	apList := make([]AttrParam, 0)
	sv := reflect.ValueOf(t.ParamValue)
	for i := 0; i < sv.Len(); i++ {
		item := sv.Index(i).Interface().(map[string]interface{})
		for _, field := range t.ChildKeys {
			arr_key := t._ori_arr_key(i, field)
			var value interface{}
			if !utils.IsEmpty(field, item) {
				value = item[field]
			}

			ap := AttrParam{AttrName: arr_key, AttrValue: value}
			apList = append(apList, ap)
		}

	}
	return apList

}
