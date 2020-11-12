package v1

import (
	"DockerUI/app/models"
	"DockerUI/app/service"
	"DockerUI/app/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	host := c.Query("host")
	infoList := make([]interface{}, 0)
	response := &models.Response{
		Code:    http.StatusOK,
		Message: "Docker主机信息获取成功!",
		Data:    infoList,
	}
	info, err := service.Info(host)
	util.ErrHandler(err, response, http.StatusInternalServerError, "Docker主机信息获取失败", infoList)
	infoList = append(infoList, info)
	c.JSON(http.StatusOK, response)
}

func ImageList(c *gin.Context) {
	host := c.Query("host")
	infoList := make([]interface{}, 0)
	response := &models.Response{
		Code:    http.StatusOK,
		Message: "image获取成功！",
		Data:    infoList,
	}
	images, err := service.ImageList(host)
	util.ErrHandler(err, response, http.StatusInternalServerError, "image获取失败！", infoList)
	infoList = append(infoList, images)
	c.JSON(http.StatusOK, response)
}
