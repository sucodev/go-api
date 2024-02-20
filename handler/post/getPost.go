package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Post) GetPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post",
	})
}
