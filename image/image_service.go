package image

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
}

func (image *ImageService) GetInterfaceName() string {
	return dbusServiceName
}

func NewImageService(service *dbusutil.Service, cli *client.Client) *ImageService {
	imageService := ImageService{
		service: service,
		cli:     cli,
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

func (image *ImageService) GetImageList() (result string, busErr *dbus.Error) {
	ctx := context.Background()
	images, err := image.cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Println("镜像列表获取失败", err)
	}

	list, _ := json.Marshal(images)
	result = string(list)
	fmt.Println(result)
	return result, nil
}
