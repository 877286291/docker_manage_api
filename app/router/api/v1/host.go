package v1

import (
	"DockerUI/app/sdk"
	"DockerUI/app/service"
	"DockerUI/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func AddHost(c *gin.Context) {
	name := c.PostForm("name")
	url := c.PostForm("url")
	cli, err := sdk.ConnInit(url)
	if err != nil {
		log.Println(err)
		return
	}
	info, err := sdk.Info(cli)
	if err != nil {
		log.Println(err)
		return
	}
	tmpMap := util.Struct2Map(info)
	tmpMap["name"] = name
	tmpMap["url"] = url
	service.AddHostToDB(tmpMap)
}
func DeleteHost(c *gin.Context) {

}
func ModifyHost(c *gin.Context) {

}
