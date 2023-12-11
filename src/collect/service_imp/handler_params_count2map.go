package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type Count2Map struct {
	BaseHandler
}

func (pr *Count2Map) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	rd := template.GetResult()
	result, ok := rd.Data.(map[string]interface{})
	if !ok {
		return common.NotOk("结果对象非map类型，请检查数据")
	}
	result[utils.GetRenderVarName(handlerParam.Field)] = rd.GetCount()
	rd.Data = result
	return common.Ok(nil, "处理参数成功")
}
