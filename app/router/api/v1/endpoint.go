package v1

import (
	"DockerUI/app/models"
	"DockerUI/app/sdk"
	"DockerUI/app/service"
	"DockerUI/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddEndPoint(c *gin.Context) {
	response := &models.Response{
		Code:    200,
		Message: "docker主机添加成功",
		Data:    nil,
	}
	hostname := c.PostForm("hostname")
	url := c.PostForm("url")
	group := c.PostForm("group")
	cli, err := sdk.ConnInit(url)
	if err != nil {
		log.Println(err)
	}
	info, err := sdk.Info(cli)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "docker主机无法连接，请检查网络连接"
		c.JSON(http.StatusOK, response)
		return
	}
	tmpMap := util.Struct2Map(info)
	tmpMap["Hostname"] = hostname
	tmpMap["Url"] = url
	tmpMap["Group"] = group
	tmpMap["SystemStatus"] = "UP"
	insertId, err := service.AddEndPointToDB(tmpMap)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "docker主机信息写入数据库失败"
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = insertId
	c.JSON(http.StatusOK, response)
}
func DeleteEndPoint(c *gin.Context) {

}
func ModifyEndPoint(c *gin.Context) {

}
