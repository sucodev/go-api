package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Post) CreatePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post",
	})
}
