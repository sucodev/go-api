package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *PostController) CreatePostController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post",
	})
}
