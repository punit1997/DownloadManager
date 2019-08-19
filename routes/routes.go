package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "working",
		})
	})

	return route
}
