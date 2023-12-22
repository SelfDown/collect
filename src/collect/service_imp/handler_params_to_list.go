package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

type ToList struct {
	BaseHandler
}

func (uf *ToList) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	// 转树形结构
	target := treeToList(arr, childName)
	r := common.Ok(target, "处理参数成功")
	return r
}

func treeToList(arr []map[string]interface{}, children string) []map[string]interface{} {
	r := make([]map[string]interface{}, 0)
	for _, aVal := range arr {
		r = append(r, aVal)
		if _, hasChildren := aVal[children]; hasChildren {
			subTreeList, _ := utils.RenderVarToArrMap(children, aVal)
			subList := treeToList(subTreeList, children)
			r = append(r, subList...)
		}

	}
	return r
}
