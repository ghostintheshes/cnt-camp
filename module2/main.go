package main

import (
	"cnt-camp/module2/app"
	"github.com/gin-gonic/gin"
)

func main() {
	Core := gin.New()
	Core.Use(gin.Logger(), gin.Recovery())
	Core.GET("/q1", app.Q1)
	Core.GET("/q2", app.Q2)
	Core.GET("/q3", app.Q3)
	Core.GET("/healthz", app.Healthz)
	Core.Run(":8080")
}
