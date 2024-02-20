package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *PostController) DeletePostController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete post",
	})
}
