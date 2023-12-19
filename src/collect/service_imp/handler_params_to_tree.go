package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
)

type ToTree struct {
	BaseHandler
}

func (uf *ToTree) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	pid := handlerParam.Pid
	id := handlerParam.Id
	// 转树形结构
	target := listToTree(arr, id, pid, childName)
	r := common.Ok(target, "处理参数成功")
	return r
}

func listToTree(arr []map[string]interface{}, id string, pid string, children string) []map[string]interface{} {
	r := make([]map[string]interface{}, 0)
	hash := make(map[interface{}]map[string]interface{})
	for _, jsonItem := range arr {
		hash[jsonItem[id]] = jsonItem
	}
	for _, aVal := range arr {
		parentId := aVal[pid]
		if hashVp, ok := hash[parentId]; ok {

			if _, hasKey := hashVp[children]; !hasKey {
				hashVp[children] = make([]map[string]interface{}, 0)
			}
			ch, _ := utils.RenderVarToArrMap(children, hashVp)
			ch = append(ch, aVal)
			hashVp[children] = ch
		} else {
			r = append(r, aVal)
		}
	}
	return r
}
