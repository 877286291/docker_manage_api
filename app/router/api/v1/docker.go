package v1

import (
	"DockerUI/app/models"
	"DockerUI/app/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Info(c *gin.Context) {
	host := c.Query("host")
	info, err := service.Info(host)
	if err != nil {
		log.Println(err)
	}
	infoList := make([]interface{}, 0)
	response := models.Response{
		Code:    0,
		Message: "",
		Data:    nil,
	}
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Docker主机信息获取失败！"
	} else {
		infoList = append(infoList, info)
		response.Code = http.StatusOK
		response.Message = "Docker主机信息获取成功！"
		response.Data = infoList
	}
	c.JSON(http.StatusOK, response)
}

func ImageList(c *gin.Context) {
	host := c.Query("host")
	images, err := service.ImageList(host)
	infoList := make([]interface{}, 1)
	response := models.Response{
		Code:    0,
		Message: "",
		Data:    nil,
	}
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "image获取失败！"
	} else {
		infoList = append(infoList, images)
		response.Code = http.StatusOK
		response.Message = "image获取成功！"
		response.Data = infoList
	}
	c.JSON(http.StatusOK, response)
}
