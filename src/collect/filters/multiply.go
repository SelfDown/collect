package collect

import "strconv"

func Multiply(a, b interface{}) string {
	var floatA, floatB float64

	// 将参数转换为 float64
	switch v := a.(type) {
	case int:
		floatA = float64(v)
	case float64:
		floatA = v
	default:
		return ""
	}

	switch v := b.(type) {
	case int:
		floatB = float64(v)
	case float64:
		floatB = v
	default:
		return ""
	}
	// 执行乘法操作
	result := floatA * floatB
	// 将结果转换为字符串
	return strconv.FormatFloat(result, 'f', 2, 64)
}
