package main

import (
	template_service "collect.mod/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	//req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	//fmt.Println(err)
	//
	//req.Header.Set("xx", "x")
	//resp, err := http.DefaultClient.Do(req)
	//fmt.Println(resp)
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
	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")
	template_service.RunScheduleService()
	r.POST("/template_data/data", func(c *gin.Context) {
		template_service.HandlerRequest(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
