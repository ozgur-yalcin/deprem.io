package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Anasayfa(ctx *gin.Context) {
	ctx.String(http.StatusOK, "deprem.io backend")
}
