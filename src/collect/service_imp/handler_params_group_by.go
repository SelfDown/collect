package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"strings"
)

type GroupBy struct {
	BaseHandler
}

func (uf *GroupBy) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	arrDict := make(map[string]map[string]interface{})
	fields := uf.GetFieldNames(handlerParam, params)
	appendItemParam := handlerParam.AppendItemParam
	orderList := make([]map[string]interface{}, 0)
	for _, item := range arr {
		valueList := utils.GetFieldValueList(fields, item)
		key := strings.Join(valueList, "_$[#]$_")
		dictItem, ok := arrDict[key]
		if !ok {
			children := make([]map[string]interface{}, 0)
			var targetMap map[string]interface{}
			if appendItemParam { // 如果添加item 的所有参数,取第一个
				targetMap = utils.CopyMap(item)
			} else { // 只添加分组的字段
				targetMap = make(map[string]interface{})
				for _, fieldName := range fields {
					field := utils.GetRenderVarName(fieldName)
					targetMap[field] = utils.RenderVar(fieldName, item)
				}
			}

			children = append(children, item)
			targetMap[childName] = children
			arrDict[key] = targetMap
			orderList = append(orderList, targetMap)
		} else {
			children := dictItem[childName].([]map[string]interface{})
			children = append(children, item)
			dictItem[childName] = children
		}

	}

	r := common.Ok(orderList, "处理参数成功")
	return r
}
