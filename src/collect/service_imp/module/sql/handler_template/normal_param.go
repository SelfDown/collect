package collect

type NormalParam struct {
	// 转换后key
	ParamKey string
	// 原始key
	ParamKeyOriginal string
	// 值
	ParamValue interface{}
	// 子变量
	ChildKeys []string
}

const SPLIT = "________________COLLECT_SQL_NORMAL_PARAM_SPLIT________________"

func (t *NormalParam) _ori_key() string {
	return "_" + t.ParamKey + SPLIT
}

func (t *NormalParam) SetChildKeys(childrenKeys []string) {
	t.ChildKeys = childrenKeys
}

func (t *NormalParam) SetParamKeyOrinal(param_key_original string) {
	// 设置原始key
	t.ParamKeyOriginal = param_key_original
	// 设置里面的变量
	param_key := GetSqlParamKey(param_key_original)
	t.SetParamKey(param_key)

}

func (t *NormalParam) SetParamKey(param_key string) {
	t.ParamKey = param_key
}

func (t *NormalParam) SetParamKeyValue(param_key_value interface{}) {
	t.ParamValue = param_key_value
}

// func (t *NormalParam) GetParamKeyValue() interface{} {
// 	return t.ParamValue
// }

/**
* to_param_key 渲染第一次模板变量为true
**/
func (t *NormalParam) GetSqlParamKeyValue(to_param_key bool) interface{} {
	if to_param_key { // 重新命名个变量
		return "{{." + t._ori_key() + "}}"
	} else {
		param_key_value := "?"
		return param_key_value
	}

}
func (t *NormalParam) GetSqlParamParamKey() string {
	return t.ParamKey
}

func (t *NormalParam) GetParamKeyOriginal() string {
	return t.ParamKeyOriginal
}

/**
* 获取变量的实际值
 */
func (t *NormalParam) GetValue() []interface{} {
	result := make([]interface{}, 0)
	result = append(result, t.ParamValue)
	return result
}

/**
* 获取数组对象的key
 */
func (t *NormalParam) GetSqlParamParamKeyList() []AttrParam {
	ap := AttrParam{AttrName: t._ori_key(), AttrValue: t.ParamValue}
	apList := make([]AttrParam, 0)
	apList = append(apList, ap)
	return apList
}
