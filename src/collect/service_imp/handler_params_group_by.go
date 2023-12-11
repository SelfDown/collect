package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
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
	target := make([]map[string]interface{}, 0)
	fields := uf.GetFieldNames(handlerParam)
	for _, item := range arr {
		valueList := utils.GetFieldValueList(fields, item)
		key := strings.Join(valueList, "_$[#]$_")
		dictItem, ok := arrDict[key]
		if !ok {
			children := make([]map[string]interface{}, 0)
			targetMap := utils.CopyMap(item)
			children = append(children, item)
			targetMap[childName] = children
			arrDict[key] = targetMap
		} else {
			children := dictItem[childName].([]map[string]interface{})
			children = append(children, item)
			dictItem[childName] = children
		}

	}
	for _, data := range arrDict {
		target = append(target, data)
	}

	r := common.Ok(target, "处理参数成功")
	return r
}
