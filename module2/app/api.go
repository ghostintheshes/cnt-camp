package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func Q1(ctx *gin.Context) {
	Header := ctx.Request.Header
	for k, v := range Header {
		for _, value := range v {
			ctx.Writer.Header().Add(k, value)
		}
	}
	ctx.JSON(200, fmt.Sprintf("Headers has been written:\n%s", Header))
}

func Q2(ctx *gin.Context) {
	os.Setenv("VERSION", "v1.0")
	ver := os.Getenv("VERSION")
	if ver != "" {
		ctx.Header("VERSION", ver)
	}
	ctx.JSON(200, map[string]string{"VERSION": ver})
}

func Q3(ctx *gin.Context) {
	Ip := ctx.ClientIP()
	//Ip := getRealIP(ctx.Request)
	log.Printf("Client IP is %s,  Return Code: %d", Ip, 200)
	ctx.Status(200)
}

func getRealIP(request *http.Request) string {
	ip := ""
	if h := request.Header.Get("X-Forwarded-For"); h != nil {
		ip = strings.TrimSpace(strings.Split(h, ",")[0])
		if ip != nil {
			return ip
		}
	}
	if h := request.Header.Get("X-Real-Ip"); h != nil {
		ip := strings.TrimSpace(strings.Split(h, ",")[0])
		if ip != nil {
			return ip
		}
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(request.RemoteAddr)); err == nil {
		return ip
	}
	return ip
}

func Healthz(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
