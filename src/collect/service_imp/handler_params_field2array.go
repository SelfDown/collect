package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"strings"
)

type Field2Array struct {
	BaseHandler
}

func (uf *Field2Array) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	field := utils.RenderVar(handlerParam.Field, params).(string)
	arr := strings.Split(field, ",")
	r := common.Ok(arr, "处理参数成功")
	return r
}
