package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type Param2Result struct {
	BaseHandler
}

func (pr *Param2Result) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()

	result, ok := params[utils.GetRenderVarName(handlerParam.Field)]
	if !ok {
		return common.NotOk(handlerParam.Field + "在参数中没有找到")
	}
	var rd *common.Result
	if template.HasResult() {
		rd = template.GetResult()
	} else {
		rd = common.Ok(nil, "成功")
	}
	rd.Data = result
	template.SetResult(rd)
	r := common.Ok(nil, "处理参数成功")
	return r
}
