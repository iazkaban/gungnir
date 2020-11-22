package routers

import "github.com/gin-gonic/gin"

func authRouter(router *gin.Engine) {
	//授权相关
	auth := router.Group("/auth")
	auth.POST("/login") //登入
	auth.GET("/logout") //登出

	router.GET("/users")          //用户列表
	user := router.Group("/user") //用户组
	user.POST("")                 //添加/修改用户
	user.GET("")                  //获取用户
	user.DELETE("")               //删除用户
}
