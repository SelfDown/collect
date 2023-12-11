package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type SessionAdd struct {
	BaseHandler
}

func (sa *SessionAdd) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	session := ts.GetSession()

	for _, field := range handlerParam.Fields {
		key := field.Key // key 作为存储字段，field 作为转参数字段
		if utils.IsValueEmpty(key) {
			return common.NotOk("session 添加器未设置 key")
		}
		if utils.IsValueEmpty(field.Template) {
			return common.NotOk("session 添加器未设置 template")
		}
		value := utils.RenderTplDataWithType(field.TemplateTpl, params, field.Type)
		(*session).Set(key, value)
	}
	(*session).Save()

	r := common.Ok(nil, "处理参数成功")
	return r
}
