package v1

import (
	"DockerUI/app/models"
	"DockerUI/app/sdk"
	"DockerUI/app/service"
	"DockerUI/app/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddHost(c *gin.Context) {
	response := &models.Response{
		Code:    200,
		Message: "docker主机添加成功",
		Data:    nil,
	}
	hostname := c.PostForm("hostname")
	url := c.PostForm("url")
	cli, err := sdk.ConnInit(url)
	util.ErrHandler(err, response, 500, "docker主机无法连接，请检查网络连接", nil)
	info, err := sdk.Info(cli)
	util.ErrHandler(err, response, 500, "docker主机信息获取失败", nil)
	tmpMap := util.Struct2Map(info)
	tmpMap["hostname"] = hostname
	tmpMap["url"] = url
	insertId, err := service.AddHostToDB(tmpMap)
	util.ErrHandler(err, response, 500, "docker主机信息写入数据库失败", nil)
	response.Data = insertId
	c.JSON(http.StatusOK, response)
}
func DeleteHost(c *gin.Context) {

}
func ModifyHost(c *gin.Context) {

}
