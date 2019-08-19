package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/punit1997/DownloadManager/download"
)

func InitRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	route.POST("/downloads", download.Start)

	return route
}
