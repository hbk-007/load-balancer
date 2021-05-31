package main

import (
	"load-balancer/cron"
	"load-balancer/pool"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	InitRoutes(router)
	pool.InitServerPool()
	cron.UpdateHealthCron()
	router.Run(":8800")
}
