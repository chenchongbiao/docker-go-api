package volume

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bluesky/docker-go-api/adapter"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

const (
	dbusPath        = "/com/bluesky/docker/Volume"
	dbusServiceName = "com.bluesky.docker.Volume"
	dbusInterface   = dbusServiceName
)

type VolumeService struct {
	service *dbusutil.Service
	cli     *client.Client
	adapter *adapter.VolumeAdapter
}

func (v *VolumeService) GetInterfaceName() string {
	return dbusServiceName
}

func NewVolumeService(service *dbusutil.Service, cli *client.Client) *VolumeService {
	volumeService := VolumeService{
		service: service,
		cli:     cli,
		adapter: adapter.NewVolumeAdapter(cli),
	}
	err := service.Export(dbusPath, &volumeService)
	if err != nil {
		log.Println(err)
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		log.Println(err)
	}
	return &volumeService
}

func (v *VolumeService) GetVolumeList(args map[string]interface{}) (result string, busErr *dbus.Error) {
	var options volume.ListOptions
	if len(args) == 0 {
		options = volume.ListOptions{}
	} else {
		filter := filters.NewArgs()
		for k := range args {
			// fmt.Printf("%#v\n", k)
			filter.Add(k, args[k].(string))
		}
		options = volume.ListOptions{Filters: filter}
	}

	volumes, err := v.cli.VolumeList(context.Background(), options)
	if err != nil {
		log.Println("存储列表获取失败", err)
		return result, nil
	}
	log.Println("存储列表获取成功")

	items := v.adapter.List(volumes)
	resultMap := map[string]interface{}{
		"status": true,
		"data":   items,
	}

	// 将map转换为json数据
	resultJson, _ := json.Marshal(resultMap)
	result = string(resultJson)

	// fmt.Printf("%#v\n", result)

	return result, nil
}
