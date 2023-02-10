package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ozgur-soft/deprem.io/cache"
)

func Flushall(ctx *gin.Context) {
	cache.Flush(ctx)
	ctx.String(http.StatusOK, "OK")
}
