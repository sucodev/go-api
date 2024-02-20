package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct{}

func (p *Post) GetPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get posts",
	})
}

func (p *Post) GetPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post",
	})
}

func (p *Post) CreatePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post",
	})
}

func (p *Post) DeletePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete post",
	})
}

func (p *Post) UpdatePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update post",
	})
}
