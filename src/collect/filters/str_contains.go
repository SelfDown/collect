package collect

import "strings"

// 定义一个函数来检查字符串是否在字符串中
func StrContains(s, substr string) bool {
	return strings.Contains(s, substr)
}
