package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
