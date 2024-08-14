package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"strings"
)

type Field2Array struct {
	BaseHandler
}

func (uf *Field2Array) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	if !utils.IsValueEmpty(handlerParam.Foreach) {
		// 直接渲染变量
		dataList, errMsg := utils.RenderVarToArrMap(handlerParam.Foreach, params)
		if !utils.IsValueEmpty(errMsg) {
			return common.NotOk(errMsg)
		}

		for _, field := range handlerParam.Fields {

			for _, item := range dataList {
				forEachField := utils.RenderVar(field.Field, item)
				if !utils.IsValueEmpty(forEachField) {
					arr := strings.Split(forEachField.(string), ",")
					item[field.SaveField] = arr
				} else {
					item[field.SaveField] = make([]string, 0)
				}

			}
		}
		r := common.Ok(nil, "处理参数成功")
		return r
	} else {
		field := utils.RenderVar(handlerParam.Field, params)
		arr := make([]string, 0)
		if !utils.IsValueEmpty(field) {
			arr = strings.Split(field.(string), ",")
		}
		r := common.Ok(arr, "处理参数成功")
		return r
	}

}
