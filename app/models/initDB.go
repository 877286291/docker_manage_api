package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 连接设置
func SetConnect() *mongo.Client {
	uri := "mongodb://127.0.0.1:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println(err)
	}
	return client
}
