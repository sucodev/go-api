package user

import "github.com/gin-gonic/gin"

type UserController struct{}

func InitializeUserRouter(api *gin.RouterGroup) {
	user := UserController{}
	{
		api.GET("/user", user.Create)
	}
}
