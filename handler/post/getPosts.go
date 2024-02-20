package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Post) GetPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get posts",
	})
}
