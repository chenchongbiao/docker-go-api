package adapter

import (
	"encoding/json"
	"strings"

	"github.com/bluesky/docker-go-api/convertor"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type VolumeAdapter struct {
	cli       *client.Client
	convertor *convertor.VolumeConvertor
}

func NewVolumeAdapter(cli *client.Client) *VolumeAdapter {
	volAdapter := VolumeAdapter{
		cli: cli,
	}
	return &volAdapter
}

func (v *VolumeAdapter) List(volumes volume.ListResponse) []map[string]interface{} {
	var respMap map[string]interface{} // 存放容器信息的map

	respJson, _ := json.Marshal(volumes) // 将容器的结构体数据转换为json数据
	json.NewDecoder(strings.NewReader(string(respJson))).Decode(&respMap)

	volumeList := respMap["Volumes"].([]interface{})
	// fmt.Printf("%#v\n", volumeList)
	items := make([]map[string]interface{}, 0, len(volumeList))
	for _, volume := range volumeList {
		item := v.convertor.VolumeConvert(volume.(map[string]interface{}), true)
		// fmt.Printf("%#v\n", volume)
		items = append(items, item)
	}
	// fmt.Printf("%#v\n", items)
	return items
}
