package main

import "collect.mod/simple"

func main() {
	//simple.QueryUserList()
	//simple.UserCreate()
	//根据用户名更新 1个条件的
	//simple.UserUpdate()
	// 根据用户名和昵称更新，2个条件的
	//simple.UserUpdateByUsernameNick()
	// 根据用户ID更新
	//simple.UserUpdateByUseridList()
	// 更新所有数据
	//simple.UserUpdateAll()
	// 删除用户
	simple.UserDeleteByUseridList()
}
