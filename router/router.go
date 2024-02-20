package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/controller"
)

func Initialize() {
	// Declaration Router
	router := gin.Default()

	// Create Group Base API
	api := router.Group("/api")

	controller.IntializeHandlers(api)

	router.Run(":3333")
}
