package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type Result2Params struct {
	BaseHandler
}

func (pr *Result2Params) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	rd := template.GetResult()
	result, ok := rd.Data.(map[string]interface{})
	if !ok {
		return common.Ok(nil, "处理参数成功")
	}
	for _, field := range handlerParam.Fields {
		if utils.IsValueEmpty(field.From) {
			return common.NotOk("结果转参数处理器中，未配置from 字段")
		}
		fromValue := utils.RenderTplDataWithType(field.FromTpl, result, field.Type)
		template.AddParam(utils.GetRenderVarName(field.To), fromValue)
	}
	return common.Ok(nil, "处理参数成功")
}
