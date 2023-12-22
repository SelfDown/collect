package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type FilterArr struct {
	BaseHandler
}

func (uf *FilterArr) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	target := make([]map[string]interface{}, 0)
	old, exists := params[handlerParam.Item]

	for _, item := range arr {
		params[handlerParam.Item] = item

		ok := utils.RenderTplDataBool(handlerParam.IfTemplateTpl, params)
		if ok {
			//复制原有的数组出来
			targetItem := utils.Copy(item).(map[string]interface{})
			target = append(target, targetItem)
		}

	}
	if exists { // 还原参数
		params[handlerParam.Item] = old
	} else {
		delete(params, handlerParam.Item)
	}

	r := common.Ok(target, "处理参数成功")
	return r
}
