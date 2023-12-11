package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type IgnoreData struct {
	BaseHandler
}

//删除数组数据
func (uf *IgnoreData) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	params := template.GetParams()
	dataList := utils.RenderVar(handlerParam.Foreach, params).([]map[string]interface{})
	resultList := make([]map[string]interface{}, 0)
	paramsName := handlerParam.Params
	for index, dataItem := range dataList {
		if !utils.IsValueEmpty(paramsName) {
			dataItem[paramsName] = params
		}
		for _, field := range handlerParam.Fields {

			if !utils.RenderTplDataBool(field.TemplateTpl, dataItem) {
				if !utils.IsValueEmpty(paramsName) {
					delete(dataItem, paramsName)
				}
				resultList = append(resultList, dataList[index])
			}
		}
	}
	// 重新设置对象
	template.AddParam(utils.GetRenderVarName(handlerParam.Foreach), resultList)
	r := common.Ok(nil, "处理参数成功")
	return r
}
