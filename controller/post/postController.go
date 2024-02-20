package post

import (
	"github.com/gin-gonic/gin"
)

type PostController struct{}

func InitializePostRouter(api *gin.RouterGroup) {
	post := PostController{}

	{
		api.GET("/posts", post.GetPostController)
		api.GET("/post", post.GetPostsController)
		api.POST("/post", post.CreatePostController)
		api.DELETE("/post", post.DeletePostController)
		api.PUT("/post", post.UpdatePostController)
	}
}
