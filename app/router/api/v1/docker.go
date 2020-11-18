package v1

import (
	"DockerUI/app/models"
	"DockerUI/app/service"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"sync"
)

func Info(c *gin.Context) {
	// 取到所有的endpoints
	infoList := make([]interface{}, 0)
	response := &models.Response{
		Code:    http.StatusOK,
		Message: "Docker主机信息获取成功!",
		Data:    nil,
	}
	allEndPoints := service.FindAllFromDB().([]models.EndPointBaseInfo)
	var wg sync.WaitGroup
	wg.Add(len(allEndPoints))
	for _, endPoint := range allEndPoints {
		//验证每一个endpoints主机状态
		go func(endPoint models.EndPointBaseInfo) {
			info, err := service.Info(endPoint.Url)
			//如果可以连接，修改数据库中的状态为up，否则为down
			if err != nil {
				//TODO 修改数据库主机SystemStatus状态为DOWN
				row := models.EndPointBaseInfo{}
				_ = info.(*mongo.SingleResult).Decode(&row)
				infoList = append(infoList, row)
				wg.Done()
				return
			}
			//TODO 修改数据库主机所有信息,SystemStatus=UP
			infoList = append(infoList, info.(types.Info))
			wg.Done()
		}(endPoint)
	}
	wg.Wait()
	response.Data = infoList
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
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "image获取失败！"
		c.JSON(http.StatusOK, response)
		return
	}
	infoList = append(infoList, images)
	c.JSON(http.StatusOK, response)
}
