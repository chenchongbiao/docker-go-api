package network

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
	dbusPath        = "/com/bluesky/docker/Network"
	dbusServiceName = "com.bluesky.docker.Network"
	dbusInterface   = dbusServiceName
)

type NetworkService struct {
	service *dbusutil.Service
	cli     *client.Client
	adapter *adapter.NetworkAdapter
}

func (network *NetworkService) GetInterfaceName() string {
	return dbusServiceName
}

func NewNetworkService(service *dbusutil.Service, cli *client.Client) *NetworkService {
	networkService := NetworkService{
		service: service,
		cli:     cli,
		adapter: adapter.NewNetworkAdapter(cli),
	}

	err := service.Export(dbusPath, &networkService)
	if err != nil {
		log.Println(err)
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		log.Println(err)
	}
	return &networkService
}

// 获取镜像列表
func (n *NetworkService) GetNetworkList() (result string, busErr *dbus.Error) {
	nets, err := n.cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		log.Println("网络列表获取失败", err)
	}
	log.Println("网络列表获取成功")

	items := n.adapter.List(nets)
	resultMap := map[string]interface{}{
		"status": true,
		"data":   items,
	}

	// 将map转换为json数据
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)
	// fmt.Printf("%#v\n", resultMap)

	// fmt.Printf("%#v\n", nets)
	return result, nil
}
