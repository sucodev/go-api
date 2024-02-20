package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *PostController) GetPostsController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get posts",
	})
}
