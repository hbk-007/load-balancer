package main

import (
	"load-balancer/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", handler.PingStatus)
	router.Any("/api/*pattern", handler.LoadBalancer)
}
