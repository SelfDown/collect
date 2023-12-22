package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type SessionRemove struct {
	BaseHandler
}

func (sr *SessionRemove) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	session := ts.GetSession()
	for _, field := range handlerParam.Fields {
		key := field.Key // key 作为存储字段，field 作为转参数字段
		if utils.IsValueEmpty(key) {
			return common.NotOk("session 添加器未设置 key")
		}
		(*session).Delete(key)
	}
	(*session).Save()

	r := common.Ok(nil, "处理参数成功")
	return r
}
