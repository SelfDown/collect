package model

import (
	utils "collect/src/collect/utils"
)

var modelMap map[string]interface{}
var primaryKeyMap map[string][]string

// 生成一个脚本自动填充这个
func init() {
	//todo 如果用hashmap 效率慢，可以换二叉树，目前1、200个表很快
	modelMap = make(map[string]interface{})
	primaryKeyMap = make(map[string][]string)
	addTable()
}
func GetModel(tableName string) interface{} {
	return modelMap[tableName]
}
func CloneModel(tableName string) interface{} {
	data := modelMap[tableName]
	return utils.Copy(data)
}
func GetPrimaryKey(tableName string) []string {
	return primaryKeyMap[tableName]
}
