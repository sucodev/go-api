package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sucodev/go-api/handler"
)

func Initialize() {
	// Declaration Router
	router := gin.Default()

	// Create Group Base API
	api := router.Group("/api")

	handler.IntializeHandlers(api)

	router.Run(":3333")
}
