package service

import (
	"DockerUI/app/models/mongo"
	"github.com/docker/docker/api/types"
)

func AddEndPointToDB(info map[string]interface{}) (interface{}, error) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	insertResult, err := mgo.InsertOne(info)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}
func DeleteEndPointFromDB(url string) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.FindOne("Url", url)
}
func UpdateEndPointFromDB(url string, info types.Info) {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	mgo.FindOne("Url", url)
}
func FindEndPointFromDB(url string) interface{} {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	return mgo.FindOne("Url", url)
}
func FindAllFromDB() interface{} {
	mgo := mongo.NewMgo("docker_info", "endpoints")
	return mgo.FindAll()
}
