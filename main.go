package main

import (
	"github.com/gin-gonic/gin"
	"github.com/talltslife/demo-dd/routes"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	routes.InitilizeBarRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
