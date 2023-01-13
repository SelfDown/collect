package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type UpdateArray struct {
	BaseHandler
}

func (uf *UpdateArray) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	// 直接渲染变量
	dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	if !utils.IsValueEmpty(errMsg) {
		return common.NotOk(errMsg)
	}
	var paramsCopy map[string]interface{}
	if !utils.IsValueEmpty(handlerParam.Item) { // 如果没有配置item 则取本身
		paramsCopy = utils.CopyMap(params)
	}
	for _, field := range handlerParam.Fields {

		for _, item := range dataList {
			if !utils.IsValueEmpty(handlerParam.Item) { // 如果配置了item，设置item
				paramsCopy[utils.ItemName] = item
			} else { // 没有配置item取整个item
				paramsCopy = item
			}
			//渲染值
			value := utils.RenderTplDataWithType(field.TemplateTpl, paramsCopy, field.Type)
			item[field.Field] = value
		}
	}
	if template.Log {
		template.LogData("update_array 处理结果")
		template.LogData(params)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
