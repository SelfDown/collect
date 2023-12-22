package main

import (
	gen "github.com/SelfDown/collect/gen"
	"github.com/SelfDown/collect/model"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main1() {
	gen.GenModel()
}
func main() {

	// todo go profile 使用
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 生成cookies
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_id", store))
	r.Static("/static", "./static")
	// 设置数据库表
	templateService.SetDatabaseModel(&model.TableData{})
	// 添加定时任务
	templateService.RunScheduleService()
	// 添加启动服务
	templateService.RunStartupService()
	r.POST("/template_data/data", func(c *gin.Context) {
		templateService.HandlerRequest(c)
	})
	r.Run(":8088") // listen and serve on 0.0.0.0:8080
}
