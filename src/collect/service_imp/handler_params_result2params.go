package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type Result2Params struct {
	BaseHandler
}

func (pr *Result2Params) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	rd := template.GetResult()
	result, ok := rd.Data.(map[string]interface{})
	if !ok { // 如果是数组形式，则直接返回

		for _, field := range handlerParam.Fields {
			template.AddParam(utils.GetRenderVarName(field.To), rd.Data)
		}
	} else { //如果是map形式则取map中对应值
		for _, field := range handlerParam.Fields {
			if utils.IsValueEmpty(field.From) { //如果没有from字段，结果直接转参数
				template.AddParam(utils.GetRenderVarName(field.To), result)
			} else {
				fromValue := utils.RenderTplDataWithType(field.FromTpl, result, field.Type)
				template.AddParam(utils.GetRenderVarName(field.To), fromValue)
			}

		}
	}

	return common.Ok(nil, "处理参数成功")
}
