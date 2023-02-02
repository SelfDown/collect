package collect

/*
SubStr 针对负数和0 做了特殊处理
1. content[0:4] ，从开始节点到第5个字符
2. content[0:0] ，从开始节点到最后1个字符
3. content[0:-1] ，从开始节点到倒数第1个字符
3. content[-6:-1] ，从倒数第6个字符到倒数第1个字符


*/
func SubStr(content string, start int, end int) string {
	if end <= 0 {
		end = len(content) + end
	}
	if start < 0 {
		start = len(content) + start
	}

	return content[start:end]

}
