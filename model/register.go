package model

var modelMap map[string]interface{}

func init() {
	//todo 如果用hashmap 效率慢，可以换二叉树，目前1、200个表很快
	modelMap = make(map[string]interface{})
	modelMap["user_account"] = UserAccount{}
}
func GetModel(tableName string) interface{} {
	return modelMap[tableName]
}
