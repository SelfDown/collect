package main

import (
	template_service "collect.mod/src/collect/service_imp"
	"github.com/gin-gonic/gin"
)

func main() {
	// 查询用户
	//simple.QueryUserList()
	// 创建用户
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
	//simple.UserDeleteByUseridList()
	// 批量创建用户
	//simple.UserCreateBulk()
	//批量更新用户
	//simple.UserUpdateBulk()
	// 模块测试

	r := gin.Default()
	r.Static("/static", "./static")
	r.POST("/template_data/data", func(c *gin.Context) {
		params := make(map[string]interface{})
		c.Bind(&params)
		ts := template_service.TemplateService{OpUser: "zhangzhi"}
		data := ts.Result(params, true)
		c.JSON(200, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
