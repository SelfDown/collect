package collect

import "strings"

func Replace(source string, from string, to string) string {
	return strings.ReplaceAll(source, from, to)
}
