package image

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bluesky/docker-go-api/adapter"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

const (
	dbusPath        = "/com/bluesky/docker/Image"
	dbusServiceName = "com.bluesky.docker.Image"
	dbusInterface   = dbusServiceName
)

type ImageService struct {
	service *dbusutil.Service
	cli     *client.Client
	adapter *adapter.ImageAdapter
}

func (image *ImageService) GetInterfaceName() string {
	return dbusServiceName
}

func NewImageService(service *dbusutil.Service, cli *client.Client) *ImageService {
	imageService := ImageService{
		service: service,
		cli:     cli,
		adapter: adapter.NewImageAdapter(cli),
	}

	err := service.Export(dbusPath, &imageService)
	if err != nil {
		log.Println(err)
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		log.Println(err)
	}
	return &imageService
}

func (i *ImageService) GetImageList() (result string, busErr *dbus.Error) {
	images, err := i.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Println("镜像列表获取失败", err)
	}
	items := i.adapter.List(images)
	resultMap := map[string]interface{}{
		"status": true,
		"data":   items,
	}
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)
	// fmt.Printf("%#v\n", resultMap)

	return result, nil
}
