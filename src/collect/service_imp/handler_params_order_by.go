package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"sort"
)

type OrderBy struct {
	BaseHandler
}

func (uf *OrderBy) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	fields := handlerParam.Fields
	sort.Slice(arr, func(i, j int) bool {
		less := true
		for _, field := range fields {
			fieldName := field.Field
			x := arr[i]
			xValue := utils.RenderVar(fieldName, x)
			y := arr[j]
			yValue := utils.RenderVar(fieldName, y)
			if xValue == yValue {
				continue
			}
			less = utils.Strval(xValue) < utils.Strval(yValue)
			rule := field.Rule
			if rule == "desc" {
				less = !less
			}
			break

		}
		return less
	})

	r := common.Ok(arr, "处理参数成功")
	return r
}
