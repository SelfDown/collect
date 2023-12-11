package main

import (
	gen "collect/gen"
	"collect/model"
	templateService "collect/src/collect/service_imp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main1() {
	gen.GenModel()
}
func main() {

	// todo go profile 使用
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
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
	r.Run() // listen and serve on 0.0.0.0:8080
}
