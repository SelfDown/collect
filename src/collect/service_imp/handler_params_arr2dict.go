package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type Arr2Dict struct {
	BaseHandler
}

func getResultDict(foreach []map[string]interface{}, field string, value string) map[string]interface{} {
	result := make(map[string]interface{})
	//foreach, ok := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	//field := handlerParam.Field
	//value := handlerParam.Value
	//if !utils.IsValueEmpty(ok) {
	//	return common.NotOk(ok)
	//}
	//循环组成字典
	for _, item := range foreach {
		fieldName := utils.RenderVar(field, item).(string)
		result[fieldName] = utils.RenderVar(value, item)
	}
	return result
}
func (ao *Arr2Dict) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	params := template.GetParams()

	foreach, ok := utils.RenderVarToArrMap(handlerParam.Foreach, params)

	if !utils.IsValueEmpty(ok) {
		return common.NotOk(ok)
	}
	field := handlerParam.Field
	value := handlerParam.Value
	if utils.IsValueEmpty(handlerParam.Children) { // 如果是一级数组
		result := getResultDict(foreach, field, value)
		r := common.Ok(result, "处理参数成功")
		return r
	} else { //如果是二级数组
		for index, item := range foreach {
			subArr, _ := utils.RenderVarToArrMap(handlerParam.Children, item)
			result := getResultDict(subArr, field, value)
			foreach[index][handlerParam.ResultName] = result
		}
		r := common.Ok(foreach, "处理参数成功")
		return r
	}

}
