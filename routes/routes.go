package router

import (
	"net/http"
	"webesapp/controllers"
	"webesapp/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.LoadHTMLGlob("templates/*")

	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Simple Web Page",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//r.GET("/GetV6NginxQurey/:esname", controllers.GetV6NginxQurey)
	r.POST("/PostV6Qurey/", controllers.PostV6Qurey)
	r.POST("/PostV6CreateQurey/", controllers.PostV6CreateQurey)
	r.POST("/CreateQurey/", controllers.CreateQurey)
	r.POST("/DeleteQurey/", controllers.DeleteQurey)
	r.GET("/SqlV6QureyHeadr/", controllers.SqlV6Qurey)
	r.GET("/SSQureyhHeadr/", controllers.SSqlV6Qurey)
	r.PUT("/PostV6UpdateQurey/", controllers.PostV6UpdateQurey)
	r.PUT("/UpdateQurey/", controllers.UpdateQurey)

	r.POST("/PostV7Qurey/", controllers.PostV7Qurey)
	r.POST("/PostV7CreateQurey/", controllers.PostV7CreateQurey)
	r.GET("/SqlV7QureyHeadr/", controllers.SqlV7Qurey)

	//r.POST("/PostV7NginxQurey/", controllers.PostV7NginxQurey)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r

}
