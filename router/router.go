package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	app := gin.Default()

	app.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	app.Run()
}
