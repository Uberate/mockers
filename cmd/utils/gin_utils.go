package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(ctx *gin.Context, err error) {
	JSON(ctx, http.StatusBadRequest, err.Error())
}

func Success(ctx *gin.Context, body any) {
	JSON(ctx, http.StatusOK, body)
}

func JSON(ctx *gin.Context, httpCode int, body any) {
	ctx.JSON(httpCode, body)
}
