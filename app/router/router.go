package router

import (
	v1 "DockerUI/app/router/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(cors.Default())
	r.GET("/index", v1.Index)
	apiV1 := r.Group("/api/v1")
	{
		// User API
		apiV1.POST("/user/login", v1.Login)
		apiV1.DELETE("/user/logout", v1.Logout)
		// Host API
		apiV1.POST("/host/addHost", v1.AddHost)
		apiV1.DELETE("/host/deleteHost", v1.DeleteHost)
		apiV1.PUT("/host/modifyHost", v1.ModifyHost)
		// Docker API
		apiV1.GET("/docker/info", v1.Info)
		apiV1.GET("/docker/imageList", v1.ImageList)
	}

	return r
}
