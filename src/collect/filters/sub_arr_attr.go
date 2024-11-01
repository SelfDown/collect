package collect

import utils "github.com/SelfDown/collect/src/collect/utils"

// 二维数组
// arr[x][field][y]['attr']
func SubArrAttr(arr []map[string]interface{}, x int, field string, y int, attr string) interface{} {
	if x >= len(arr) {
		return ""
	}
	item := arr[x]

	subArr, _ := utils.RenderVarToArrMap(field, item)
	if y >= len(subArr) {
		return ""
	}
	return subArr[y][attr]
}
