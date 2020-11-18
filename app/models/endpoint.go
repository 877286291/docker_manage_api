package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EndPointBaseInfo struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	EndPointID        string             `json:"ID" bson:"ID"`
	Url               string             `json:"Url" bson:"Url"`
	Group             string
	Hostname          string
	SystemStatus      string
	Containers        int       `json:"Containers" bson:"Containers"`
	ContainersRunning int       `json:"ContainersRunning" bson:"ContainersRunning"`
	ContainersPaused  int       `json:"ContainersPaused" bson:"ContainersPaused"`
	ContainersStopped int       `json:"ContainersStopped" bson:"ContainersStopped"`
	Images            int       `json:"Images" bson:"Images"`
	NCPU              int       `json:"NCPU" bson:"NCPU"`
	MemTotal          int       `json:"MemTotal" bson:"MemTotal"`
	Name              string    `json:"Name" bson:"Name"`
	SystemTime        time.Time `json:"SystemTime" bson:"SystemTime"`
	Plugins           map[string][]string
}
