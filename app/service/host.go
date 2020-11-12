package service

import (
	"DockerUI/app/models/mongo"
	"github.com/docker/docker/api/types"
)

func AddHostToDB(info map[string]interface{}) (interface{}, error) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	insertResult, err := mgo.InsertOne(info)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
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
