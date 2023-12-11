package collect

import utils "collect/src/collect/utils"

func Must(value interface{}) bool {
	return !utils.IsValueEmpty(value)
}
