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
	fmt.Printf("%#v\n", conMap)
	// fmt.Printf("%#v\n", conMap["Image"])

	// 根据Image获取镜像
	filter := filters.NewArgs()
	filter.Add("reference", conMap["Image"].(string))
	image, _ := c.cli.ImageList(context.Background(), types.ImageListOptions{Filters: filter})

	item := map[string]interface{}{
		"id":    conMap["Id"],
		"name":  conMap["Names"],
		"image": c.imgConvertor.ImageConvert(image[0], false),
	}
	fmt.Printf("%#v\n", item)
	if verbose == false {
		var conInspectMap map[string]interface{} // 存放容器详细信息的map
		fmt.Println(conInspectMap)
	}
	return nil
}
