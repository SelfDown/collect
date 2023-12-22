package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
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
