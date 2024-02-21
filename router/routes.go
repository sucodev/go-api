package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/controller"
)

func Initialize() {
	// Declaration Router
	router := gin.Default()

	// Initialize Controller
	controller.IntializeControllers()

	// Create Group Base API
	api := router.Group("/api")
	{
		api.POST("/post", controller.CreatePostController)
	}

	router.Run(":3333")
}
