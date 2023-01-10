package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type CheckField struct {
	BaseHandler
}

func (uf *CheckField) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	for _, field := range handlerParam.Fields {

		ok := utils.RenderTplDataBool(field.TemplateTpl, params)
		if !ok {
			errMsg := utils.RenderTpl(field.ErrMsgTpl, params)
			return common.NotOk(errMsg)
		}
	}
	r := common.Ok(nil, "检查参数成功")
	return r
}
