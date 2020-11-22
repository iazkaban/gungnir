package controllers

import (
	"gungnir/models"
	"gungnir/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddTemplate(ctx *gin.Context) {
	temp := &models.Template{}

	err := ctx.BindJSON(temp)
	if err != nil {
		Error400(ctx, "解析body失败:"+err.Error())
		return
	}

	rs, err := services.AddTemplate(temp)
	if err != nil {
		Error500(ctx, "添加模板失败:"+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ResponseBody{
		Data: rs,
	})
}

func ListTemplate(ctx *gin.Context) {
	list, err := services.ListTemplate()
	if err != nil {
		Error500(ctx, "添加模板失败:"+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, ResponseBody{
		Data: list,
	})
}

func DeleteTemplate(ctx *gin.Context) {
	id := ctx.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		Error400(ctx, "id值类型不对:"+err.Error())
		return
	}

	err = services.DeleteTemplate(int64(idNum))
	if err != nil {
		Error500(ctx, "删除失败:"+err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
