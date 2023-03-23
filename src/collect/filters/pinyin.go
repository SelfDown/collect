package collect

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func GetPinyin(txt string) string {
	sb := new(strings.Builder)

	for _, c := range txt {
		if c > 128 {
			str := string([]rune{c})
			strs := pinyin.LazyConvert(str, nil)

			if len(strs) > 0 {
				sb.WriteString(strs[0])
			} else {
				sb.WriteString(str)
			}

		} else {
			sb.WriteString(string([]rune{c}))
		}
	}

	return strings.TrimSpace(sb.String())
}
func Pinyin(source string) string {

	return GetPinyin(source)
}
