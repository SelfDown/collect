package collect

import "strings"

func Index(source string, target string) int {
	return strings.Index(source, target)
	//return strings.ReplaceAll(source, from, to)
}
