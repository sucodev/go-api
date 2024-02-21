package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/schemas"
)

func CreatePostController(ctx *gin.Context) {
	request := CreatePostRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	post := schemas.Post{
		Title:       request.Title,
		Description: request.Description,
		Public:      *request.Public,
		Company:     request.Company,
	}

	if err := db.Create(&post).Error; err != nil {
		logger.Errorf("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating post")
		return
	}

	sendSuccess(ctx, "create-post", post)

}
