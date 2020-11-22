package routers

import (
	"gungnir/controllers"

	"github.com/gin-gonic/gin"
)

func templateRouters(router *gin.Engine) {
	router.GET("/templates", controllers.ListTemplate)

	tempRouter := router.Group("/template")

	tempRouter.POST("", controllers.AddTemplate)          //增加，修改
	tempRouter.GET("/:id")                                //获取
	tempRouter.DELETE("/:id", controllers.DeleteTemplate) //删除
}
