package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
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

//func listToTree(arr []map[string]interface{}, id string, pid string, children string) []map[string]interface{} {
//	r := make([]map[string]interface{}, 0)
//	hash := make(map[interface{}]map[string]interface{})
//	parentMap := make(map[interface{}][]map[string]interface{})
//
//	// Build the hash map and parent map
//	for _, jsonItem := range arr {
//		itemID := jsonItem[id]
//		hash[itemID] = jsonItem
//		parentID := jsonItem[pid]
//		if _, ok := parentMap[parentID]; !ok {
//			parentMap[parentID] = make([]map[string]interface{}, 0)
//		}
//		parentMap[parentID] = append(parentMap[parentID], jsonItem)
//	}
//
//	// Function to recursively build the tree
//	var buildTree func(parentID interface{}) []map[string]interface{}
//	buildTree = func(parentID interface{}) []map[string]interface{} {
//		childrenList := parentMap[parentID]
//		if len(childrenList) == 0 {
//			return nil
//		}
//		tree := make([]map[string]interface{}, 0)
//		for _, child := range childrenList {
//			childID := child[id]
//			childTree := buildTree(childID)
//			if childTree != nil {
//				child[children] = childTree
//			} else {
//				child[children] = make([]map[string]interface{}, 0)
//			}
//			tree = append(tree, child)
//		}
//		return tree
//	}
//
//	// Build the tree starting from the root nodes
//	for _, jsonItem := range arr {
//		parentID := jsonItem[pid]
//		if _, ok := hash[parentID]; !ok {
//			r = append(r, jsonItem)
//		}
//	}
//
//	// Attach children to their respective parents
//	for _, root := range r {
//		rootID := root[id]
//		root[children] = buildTree(rootID)
//	}
//
//	return r
//}

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
