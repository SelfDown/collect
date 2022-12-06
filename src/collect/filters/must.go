package collect

import utils "collect.mod/src/collect/utils"

func Must(value interface{}) bool {
	return !utils.IsValueEmpty(value)
}
