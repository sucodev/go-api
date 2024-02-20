package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/handler/post"
	"github.com/sucodev/go-api/handler/user"
)

func IntializeHandlers(api *gin.RouterGroup) {
	// Make a call for the POST package handler
	post.InitializePostRouter(api)
	user.InitializeUserRouter(api)
}
