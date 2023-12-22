package collect

import utils "github.com/SelfDown/collect/src/collect/utils"

func IsEmpty(value interface{}) bool {
	return utils.IsValueEmpty(value)
}
