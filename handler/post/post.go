package post

import (
	"github.com/gin-gonic/gin"
)

type Post struct{}

func InitializePostRouter(api *gin.RouterGroup) {
	post := Post{}

	{
		api.GET("/posts", post.GetPosts)
		api.GET("/post", post.GetPosts)
		api.POST("/post", post.CreatePost)
		api.DELETE("/post", post.DeletePost)
		api.PUT("/post", post.UpdatePost)
	}
}
