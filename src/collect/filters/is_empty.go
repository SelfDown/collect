package collect

import utils "collect/src/collect/utils"

func IsEmpty(value interface{}) bool {
	return utils.IsValueEmpty(value)
}
