package sdk

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

var cxt context.Context

func Ping(host string) {

}

func ConnInit(host string) (*client.Client, error) {
	cxt = context.Background()
	cli, err := client.NewClient(host, "", nil, nil)
	if err != nil {
		// 连接失败
		log.Println(err)
	}
	//defer cli.Close()
	return cli, err
}
func Info(cli *client.Client) (types.Info, error) {
	info, err := cli.Info(cxt)
	if err != nil {
		log.Println(err)
	}
	return info, err
}
func ImageList(cli *client.Client) ([]types.ImageSummary, error) {
	images, err := cli.ImageList(cxt, types.ImageListOptions{})
	if err != nil {
		log.Println(err)
	}
	return images, err
}
