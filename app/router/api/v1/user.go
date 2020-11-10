package v1

import (
	"DockerUI/app/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	service.Login(username, password)
}
func Logout(c *gin.Context) {
	userId := c.Param("user_id")
	service.Logout(userId)
}
