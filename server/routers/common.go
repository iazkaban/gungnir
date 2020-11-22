package routers

import "github.com/gin-gonic/gin"

func Load(router *gin.Engine) {
	authRouter(router)
	templateRouters(router)
}
