package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	post := handler.Post{}

	{
		api.GET("/posts", post.GetPosts)
		api.GET("/post", post.GetPosts)
		api.POST("/post", post.CreatePost)
		api.DELETE("/post", post.DeletePost)
		api.PUT("/post", post.UpdatePost)
	}
}
