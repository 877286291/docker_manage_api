package service

import (
	"DockerUI/app/sdk"
	"github.com/docker/docker/api/types"
	"log"
)

func Info(host string) (interface{}, error) {
	cli, err := sdk.ConnInit(host)
	if err != nil {
		log.Println(err)
	}
	info, err := sdk.Info(cli)
	if err != nil {
		//连接失败，去数据库查找
		return FindEndPointFromDB(host), err
	}
	return info, nil
}

func ImageList(host string) ([]types.ImageSummary, error) {
	cli, err := sdk.ConnInit("tcp://" + host)
	if err != nil {
		log.Println(err)
	}
	images, err := sdk.ImageList(cli)
	if err != nil {
		return nil, err
	}
	return images, nil
}
