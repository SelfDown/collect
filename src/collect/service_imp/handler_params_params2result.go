package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type Params2Result struct {
	BaseHandler
}

func (pr *Params2Result) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	result := make(map[string]interface{})
	for _, field := range handlerParam.Fields {
		if utils.IsValueEmpty(field.From) {
			return common.NotOk("参数转结果处理器中，未配置from 字段")
		}
		fromValue := utils.RenderTplDataWithType(field.FromTpl, params, field.Type)
		var toField string
		// 先从参数从取一次
		if utils.IsRenderVar(field.To) {
			toField, _ = params[utils.GetRenderVarName(field.To)].(string)

		}
		// 如果取不到就取配置的字段
		if utils.IsValueEmpty(toField) {
			toField = field.To
		}

		result[toField] = fromValue
	}

	var rd *common.Result
	if template.HasResult() {
		rd = template.GetResult()
	} else {
		rd = common.Ok(nil, "成功")
	}
	rd.Data = result
	template.SetResult(rd)
	return common.Ok(nil, "处理参数成功")
}
