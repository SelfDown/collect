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
	serviceParam := utils.GetServiceParam(handlerParam.Service, params, handlerParam.AppendParam)
	if handlerParam.AppendItemParam {
		itemMap, msg := utils.RenderVarToMap(handlerParam.Item, params)
		if utils.IsValueEmpty(msg) {
			for key, value := range itemMap {
				serviceParam[key] = value
			}
		} else {
			return common.NotOk(msg)
		}

	}
	r2 := ts.ResultInner(serviceParam)
	return r2
}
