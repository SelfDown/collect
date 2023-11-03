package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

/**
* 接收数组
 */
type PropArr struct {
	BaseHandler
}

func (uf *PropArr) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	li := make([]interface{}, 0)
	for _, item := range arr {
		value := utils.RenderVar(handlerParam.Value, item)
		li = append(li, value)
	}
	return common.Ok(li, "处理成功")
}
