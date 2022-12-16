package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	service_imp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
)

type UpdateArray struct {
	service_imp.BaseHandler
}

func (uf *UpdateArray) HandlerData(template *config.Template, handlerParam *config.HandlerParam) *common.Result {
	params := template.GetParams()
	// 直接渲染变量
	dataList, ok := utils.RenderVar(handlerParam.Foreach, params).([]map[string]interface{})
	if !ok {
		return common.NotOk(handlerParam.Foreach + "不是对象数组")
	}

	for _, field := range handlerParam.Fields {

		for _, item := range dataList {
			var paramsCopy map[string]interface{}
			if utils.IsValueEmpty(handlerParam.Item) { // 如果没有配置item 则取本身
				paramsCopy = utils.CopyMap(item)
			} else { // 如果配置了item，则从params取起
				paramsCopy = utils.CopyMap(params)
				paramsCopy[utils.ItemName] = item
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
