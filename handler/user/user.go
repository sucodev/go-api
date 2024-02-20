package user

import "github.com/gin-gonic/gin"

type User struct{}

func InitializeUserRouter(api *gin.RouterGroup) {
	user := User{}
	{
		api.GET("/user", user.Create)
	}
}
