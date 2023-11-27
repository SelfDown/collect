package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type CombineArray struct {
	BaseHandler
}

func (uf *CombineArray) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	// 直接渲染变量
	dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	// 直接渲染变量
	rightList, errMsg := utils.RenderVarToArrMap(handlerParam.Right, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	leftDict := make(map[string]map[string]interface{})
	leftFieldName := handlerParam.Field
	rightFieldName := handlerParam.Field
	if !utils.IsValueEmpty(handlerParam.RightField) {
		rightFieldName = handlerParam.RightField
	}
	children := handlerParam.Children
	// 转换左边的字典
	for _, leftItem := range dataList {
		leftField := utils.Strval(utils.RenderVar(leftFieldName, leftItem))
		leftItem[children] = make([]map[string]interface{}, 0)
		leftDict[leftField] = leftItem
	}
	// 右边挨个插入到children
	for _, rightItem := range rightList {
		rightField := utils.Strval(utils.RenderVar(rightFieldName, rightItem))
		if leftItem, ok := leftDict[rightField]; ok {
			arr := leftItem[children].([]map[string]interface{})
			arr = append(arr, rightItem)
			leftItem[children] = arr
		}
	}

	if template.Log {
		template.LogData("combine_array 处理结果")
		template.LogData(params)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
