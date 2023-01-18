package convertor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ContainerConvertor struct {
	cli          *client.Client
	imgConvertor *ImageConvertor
}

func NewContainerConvertor(cli *client.Client) *ContainerConvertor {
	conConveror := ContainerConvertor{
		cli:          cli,
		imgConvertor: NewImageConvertor(cli),
	}
	return &conConveror
}

func (c *ContainerConvertor) ContainerConvert(container types.Container, verbose bool) map[string]interface{} {
	var conMap map[string]interface{} // 存放容器信息的map

	containerJson, _ := json.Marshal(container) // 将容器的结构体数据转换为json数据
	json.NewDecoder(strings.NewReader(string(containerJson))).Decode(&conMap)

	// 根据Image获取镜像
	filter := filters.NewArgs()
	filter.Add("reference", conMap["Image"].(string))
	image, _ := c.cli.ImageList(context.Background(), types.ImageListOptions{Filters: filter})

	item := map[string]interface{}{
		"id":          conMap["Id"],
		"name":        conMap["Names"],
		"image":       c.imgConvertor.ImageConvert(image[0], false),
		"create_time": conMap["Created"],
		"status":      conMap["Status"],
		"state":       conMap["State"],
	}

	if verbose == true {
		// var conInspectMap map[string]interface{} // 存放容器详细信息的map
		// conInspec, _ := c.cli.ContainerInspect(context.Background(), conMap["Id"].(string))
		// conInspecJson, _ := json.Marshal(conInspec)
		// json.NewDecoder(strings.NewReader(string(conInspecJson))).Decode(&conInspectMap)

		netSetting := conMap["NetworkSettings"].(map[string]interface{})

		// 构造一个列表存放Networks里的所有key值
		netsList := make([]string, 0, len(netSetting))
		for value := range netSetting["Networks"].(map[string]interface{}) {
			netsList = append(netsList, value)
		}

		netsMap := netSetting["Networks"].(map[string]interface{})
		nets := netsMap[netsList[0]].(map[string]interface{})

		// config :=
		// 将镜像传入专转换器获取数据
		imgItem := c.imgConvertor.ImageConvert(image[0], true)

		item["cmd"] = imgItem["cmd"]
		item["open_stdin"] = imgItem["open_stdin"]
		item["network"] = map[string]interface{}{
			"ip":          nets["IPAddress"],
			"prefix":      nets["IPPrefixLen"],
			"gateway":     nets["Gateway"],
			"mac_address": nets["MacAddress"],
			// ports=PortMappingConvertor.from_docker(host_cfg['PortBindings']),
			"ports": PortsMapConvert(conMap["Ports"].([]interface{})),
		}
		// mounts=MountsConvertor.from_docker(attrs['Mounts']),
		item["mounts"] = MountsConvert(conMap["Mounts"].([]interface{}))
		fmt.Printf("%#v", item)

	}

	return item
}
