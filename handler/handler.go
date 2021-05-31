package handler

import (
	"fmt"
	"load-balancer/models"
	"load-balancer/pool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hey I am Alive",
	})
}

func LoadBalancer(c *gin.Context) {
	serverPool := pool.GetServerPool()
	server := models.RoundRobinScheduler(serverPool)
	fmt.Println("server is: ", server)
	server.ReverseProxy.ServeHTTP(c.Writer, c.Request)
}
