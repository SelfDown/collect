package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type UpdateOrder struct {
	BaseHandler
}

func (uf *UpdateOrder) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	field := handlerParam.Field

	// 转树形结构
	updateOrder(arr, field, childName)
	r := common.Ok(nil, "处理参数成功")
	return r
}

func updateOrder(arr []map[string]interface{}, field string, children string) {

	for index, aVal := range arr {
		num := (index + 1) * 10
		aVal[field] = num
		if _, ok := aVal[children]; ok {
			subList, _ := utils.RenderVarToArrMap(children, aVal)
			if len(subList) > 0 {
				updateOrder(subList, field, children)
			}
		}

	}
}
