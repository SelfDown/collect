package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type SessionGet struct {
	BaseHandler
}

func (sa *SessionGet) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	session := ts.GetSession()
	for _, field := range handlerParam.Fields {
		key := field.Key // key 作为存储字段，field 作为转参数字段
		if utils.IsValueEmpty(key) {
			return common.NotOk("session 添加器未设置 key")
		}
		if utils.IsValueEmpty(field.Field) {
			return common.NotOk("session 添加器未设置 field")
		}
		value := (*session).Get(key)
		template.AddParam(field.Field, value)
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
