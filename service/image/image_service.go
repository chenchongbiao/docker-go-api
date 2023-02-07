package image

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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

// 获取镜像列表
func (i *ImageService) GetImageList() (result string, busErr *dbus.Error) {
	images, err := i.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Println("镜像列表获取失败", err)
	}
	log.Println("镜像列表获取成功")

	items := i.adapter.List(images)
	resultMap := map[string]interface{}{
		"status": true,
		"data":   items,
	}

	// 将map转换为json数据
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)
	// fmt.Printf("%#v\n", resultMap)

	return result, nil
}

// 根据镜像id获取镜像
func (i *ImageService) SearchImageById(id string) (result string, busErr *dbus.Error) {
	item := i.adapter.Item(id)

	// 将map转换为json数据
	resultJson, _ := json.Marshal(item)
	result = string(resultJson)
	return result, nil
}

// 拉取镜像
func (i *ImageService) PullImage(img string) (busErr *dbus.Error) {
	out, err := i.cli.ImagePull(context.Background(), img, types.ImagePullOptions{})
	if err != nil {
		log.Println("镜像拉取失败 ", err)
	}
	io.Copy(os.Stdout, out)
	log.Println("镜像拉取成功")
	return nil
}

func (i *ImageService) PullPrivateImage(img, user, password string) (busErr *dbus.Error) {
	authConfig := types.AuthConfig{
		Username: user,
		Password: password,
	}
	encodedJson, _ := json.Marshal(authConfig)
	authStr := base64.URLEncoding.EncodeToString(encodedJson)
	out, err := i.cli.ImagePull(context.Background(), img, types.ImagePullOptions{RegistryAuth: authStr})
	out.Close()
	if err != nil {
		log.Println("私有镜像拉取失败 ", err)
	}
	io.Copy(os.Stdout, out)
	// result = "镜像拉取成功"
	log.Println("私有镜像拉取成功")
	return nil
}

func (i *ImageService) SearchImage(img string) (result string, busErr *dbus.Error) {
	images, _ := i.cli.ImageSearch(context.Background(), img, types.ImageSearchOptions{})
	// 将map转换为json数据
	resultJson, _ := json.Marshal(images)
	result = string(resultJson)
	fmt.Printf("%#v", images)
	return result, nil
}
