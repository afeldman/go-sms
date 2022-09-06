package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sms", nil)
		}
	}

	router.Run()
}
