package container

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bluesky/docker-go-api/adapter"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

const (
	dbusPath        = "/com/bluesky/docker/Container"
	dbusServiceName = "com.bluesky.docker.Container"
	dbusInterface   = dbusServiceName
)

type ContainerService struct {
	service *dbusutil.Service
	cli     *client.Client
	adapter *adapter.ContainerAdapter
}

func (c *ContainerService) GetInterfaceName() string {
	return dbusServiceName
}

func NewContainerService(service *dbusutil.Service, cli *client.Client) *ContainerService {
	containerService := ContainerService{
		service: service,
		cli:     cli,
		adapter: adapter.NewContainerAdapter(cli),
	}
	err := service.Export(dbusPath, &containerService)
	if err != nil {
		log.Println(err)
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		log.Println(err)
	}
	return &containerService
}

func (c *ContainerService) GetContainerList(args map[string]interface{}) (result string, busErr *dbus.Error) {
	var options types.ContainerListOptions
	if len(args) == 0 {
		options = types.ContainerListOptions{All: true}
	} else {
		filter := filters.NewArgs()
		for k := range args {
			// fmt.Printf("%#v\n", k)
			filter.Add(k, args[k].(string))
		}
		options = types.ContainerListOptions{Filters: filter}
	}

	containers, err := c.cli.ContainerList(context.Background(), options)
	if err != nil {
		log.Println("获取容器列表失败", err)
		return result, nil
	}
	log.Println("容器列表获取成功")

	items := c.adapter.List(containers)

	resultMap := map[string]interface{}{
		"status": true,
		"data":   items,
	}

	// 将map转换为json数据
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)

	return result, nil
}

func (c *ContainerService) Item(id string) (result string, busErr *dbus.Error) {
	item := c.adapter.Item(id)
	resultMap := map[string]interface{}{
		"status": true,
		"data":   item,
	}
	// 将map转换为json数据
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)
	log.Println("容器数据获取成功")
	return result, nil
}

func (c *ContainerService) StartContainer(ids []string) (busErr *dbus.Error) {
	for _, id := range ids {
		c.adapter.Start(id)
	}
	// err := c.cli.ContainerStart(context.Background(), ids, types.ContainerStartOptions{})
	// if err != nil {
	// 	log.Println("容器启动失败 ", err.Error())
	// 	return nil
	// }
	// log.Println("容器启动成功")
	return nil
}

func (c *ContainerService) StopContainer(ids []string) (busErr *dbus.Error) {
	for _, id := range ids {
		c.adapter.Stop(id)
	}
	return nil
}

func (c *ContainerService) RestartContainer(ids []string) (busErr *dbus.Error) {
	for _, id := range ids {
		c.adapter.Restart(id)
	}
	return nil
}

func (c *ContainerService) RmContainer(id string) (busErr *dbus.Error) {
	err := c.cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		log.Println("容器删除失败 ", err.Error())
	}
	log.Println("容器删除成功")
	return nil
}
