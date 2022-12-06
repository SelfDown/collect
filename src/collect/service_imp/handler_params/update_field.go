package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	service_imp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
)

type UpdateField struct {
	service_imp.BaseHandler
}

func (uf *UpdateField) HandlerData(template *config.Template, handlerParam *config.HandlerParam) *common.Result {

	params := template.GetParams()
	for _, field := range handlerParam.Fields {

		value := utils.RenderTplDataWithType(field.TemplateTpl, params, field.Type)
		template.AddParam(field.Field, value)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
