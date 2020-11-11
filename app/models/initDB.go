package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var DB *Database

//初始化
func init() {
	DB = &Database{
		Mongo: SetConnect(),
	}
}

// 连接设置
func SetConnect() *mongo.Client {
	uri := "mongodb://39.108.180.201:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println(err)
	}
	return client
}
