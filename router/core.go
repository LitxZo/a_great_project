package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitCoreRouter() {
	RegistRouter(func(rg *gin.RouterGroup) {
		rg.GET("/feed", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "feed success",
			})
		})
		rg.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "userInfo success",
			})
		})
		userGroup := rg.Group("/user")
		userGroup.POST("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "regist success",
			})
		})
		userGroup.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "login success",
			})
		})
		publishGroup := rg.Group("/publish")
		publishGroup.POST("/action", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "publishAction success",
			})
		})
		publishGroup.GET("/list", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, gin.H{
				"msg": "publishList success",
			})
		})
	})
}
