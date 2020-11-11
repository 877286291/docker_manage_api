package service

import (
	"DockerUI/app/sdk"
	"github.com/docker/docker/api/types"
	"log"
)

func Ping(host string) {

}

func Info(host string) (types.Info, error) {
	cli, err := sdk.ConnInit("tcp://" + host)
	if err != nil {
		log.Println(err)
	}
	info, err := sdk.Info(cli)
	if err != nil {
		return types.Info{}, err
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
