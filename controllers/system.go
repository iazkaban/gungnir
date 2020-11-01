package controllers

import (
	"gungnir/models"

	"github.com/gin-gonic/gin"
)

func System(ctx *gin.Context) {
	ctx.JSON(200, models.SystemInfo{})
}
