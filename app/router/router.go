package router

import (
	v1 "DockerUI/app/router/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode("debug")
	r.Use(cors.Default())
	r.GET("/index", v1.Index)
	apiV1 := r.Group("/api/v1")
	{
		// User API
		apiV1.POST("/user/login", v1.Login)
		apiV1.DELETE("/user/logout", v1.Logout)
		// Docker API
		apiV1.GET("/docker/info", v1.Info)
		apiV1.GET("/docker/imageList", v1.ImageList)
	}

	return r
}
