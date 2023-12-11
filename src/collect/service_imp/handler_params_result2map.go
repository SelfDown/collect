package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type Result2Map struct {
	BaseHandler
}

func (pr *Result2Map) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	rd := template.GetResult()
	r := make(map[string]interface{})
	field := utils.RenderVarOrValue(handlerParam.Field, template.GetParams()).(string)
	r[field] = rd.Data
	rd.Data = r
	return common.Ok(nil, "处理参数成功")
}
