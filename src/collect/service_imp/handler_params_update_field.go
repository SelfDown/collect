package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type UpdateField struct {
	BaseHandler
}

func (uf *UpdateField) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	for _, field := range handlerParam.Fields {

		value := utils.RenderTplDataWithType(field.TemplateTpl, params, field.Type)
		template.AddParam(field.Field, value)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
