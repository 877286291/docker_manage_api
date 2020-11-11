package service

import (
	"DockerUI/app/models/mongo"
	"github.com/docker/docker/api/types"
)

func AddHostToDB(info map[string]interface{}) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.InsertOne(info)
}
func DeleteHostFromDB(url string) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.FindOne("url", url)
}
func UpdateHostFromDB(url string, info types.Info) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.FindOne("url", url)
}
func FindHostFromDB(url string) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.FindOne("url", url)
}
