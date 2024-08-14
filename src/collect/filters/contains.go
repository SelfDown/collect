package collect


// 定义一个函数来检查字符串是否在数组中
func Contains(arr []interface{}, str interface{}) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}