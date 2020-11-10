package v1

import (
	"DockerUI/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	host := c.Query("host")
	info, err := service.Info(host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, info)
	}
}

func ImageList(c *gin.Context) {
	host := c.Query("host")
	images, err := service.ImageList(host)
	infoList := make([]interface{}, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		infoList = append(infoList, images)
		c.JSON(http.StatusOK, infoList)
	}
}
