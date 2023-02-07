package collect

import (
	utils "collect.mod/src/collect/utils"
)

func GetKey(key string) string {
	return utils.GetAppKey(key)
}
