package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Q1(ctx *gin.Context) {
	Header := ctx.Request.Header
	for k, v := range Header {
		ctx.Header(k, v[0])
	}
	ctx.JSON(200, Header)
}

func Q2(ctx *gin.Context) {
	ver := os.Getenv("VERSION")
	if ver != "" {
		ctx.Header("VERSION", ver)
	}
	ctx.JSON(200, map[string]string{"VERSION": ver})
}

func Q3(ctx *gin.Context) {
	//gin提供log中间件，默认输出
	ctx.JSON(200, "logging done")
}

func Healthz(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
