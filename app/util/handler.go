package util

import (
	"DockerUI/app/models"
	"log"
)

func ErrHandler(err error, response *models.Response, code int, message string, data interface{}) {
	if err != nil {
		log.Println(err)
		response.Code = code
		response.Message = message
		response.Data = data
	}
}
