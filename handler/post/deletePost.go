package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Post) DeletePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete post",
	})
}
