package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get post",
			})
		})
		api.POST("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "publish post",
			})
		})
		api.DELETE("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete post",
			})
		})
		api.PUT("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "put post",
			})
		})
		api.GET("/posts", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get all posts",
			})
		})
	}
}
