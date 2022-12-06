package collect

import utils "collect.mod/src/collect/utils"

func IsEmpty(value interface{}) bool {
	return utils.IsValueEmpty(value)
}
