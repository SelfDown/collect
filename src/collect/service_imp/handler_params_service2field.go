package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
)

type Service2Field struct {
	BaseHandler
}

/**
* 只做了拼接参数，未做渲染
 */
func (uf *Service2Field) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()

	//构造服务参数
	serviceParam := handlerParam.Service
	for key, value := range serviceParam {
		valueStr, ok := value.(string)
		// 判断是否为，参数变量，如果参数变量直接取参数值
		if ok && utils.IsRenderVar(valueStr) {
			val := utils.RenderVar(valueStr, params)
			serviceParam[key] = val
		}
	}
	//拼接剩余参数
	if handlerParam.AppendParam {
		for key, value := range params {
			if _, ok := serviceParam[key]; !ok {
				serviceParam[key] = value
			}
		}
	}
	r2 := ts.ResultInner(serviceParam)
	return r2
}
