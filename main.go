package main

import (
	template_service "collect.mod/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")
	// 添加定时任务
	template_service.RunScheduleService()
	// 添加启动服务
	template_service.RunStartupService()
	r.POST("/template_data/data", func(c *gin.Context) {
		template_service.HandlerRequest(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
