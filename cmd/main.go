package main

import (
	"github.com/gin-gonic/gin"
	"spydelic-api/internal/handlers"
)

// TODO: create service storage
func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", handlers.Root)
		v1.POST("/v1/create", handlers.CreateConfig)
		v1.POST("/v1/create", handlers.DeleteConfig)
	}

	router.Run(":3000")
}
