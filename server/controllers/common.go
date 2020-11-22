package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Code    int
	Message string
	Data    interface{}
}

func Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, ResponseBody{
		Code:    status,
		Message: message,
	})
}

func Error404(ctx *gin.Context, message string) {
	Error(ctx, http.StatusNotFound, message)
}

func Error400(ctx *gin.Context, message string) {
	Error(ctx, http.StatusBadRequest, message)
}

func Error403(ctx *gin.Context, message string) {
	Error(ctx, http.StatusForbidden, message)
}

func Error500(ctx *gin.Context, message string) {
	Error(ctx, http.StatusInternalServerError, message)
}
