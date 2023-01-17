package convertor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ImageConvertor struct {
	cli *client.Client
}

func NewImageConvertor(cli *client.Client) *ImageConvertor {
	imgConveror := ImageConvertor{
		cli: cli,
	}
	return &imgConveror
}

func (i *ImageConvertor) ImageConvert(image types.ImageSummary, verbose bool) map[string]interface{} {
	var imgMap map[string]interface{}        // 保存镜像信息的map
	var imgInspectMap map[string]interface{} // 保存镜像详细信息的map

	imgJson, _ := json.Marshal(image) // 镜像的ImageSummary类型转换编码成json字符串
	// _ = json.Unmarshal([]byte(imgByte), &imgMap) // 镜像的json字符串类型转换Map
	json.NewDecoder(strings.NewReader(string(imgJson))).Decode(&imgMap)

	// 获取镜像id
	id, _ := imgMap["Id"].(string)
	// 构造一个新的map，存放最终结果
	item := map[string]interface{}{
		"id":          id[7:], // 切割字符串
		"tags":        imgMap["RepoTags"],
		"author":      imgMap["Author"],
		"create_time": imgMap["Created"],
		"size":        imgMap["Size"],
	}

	if verbose == true {
		imgInspect, _, _ := i.cli.ImageInspectWithRaw(context.Background(), id)
		imgInspectJson, _ := json.Marshal(imgInspect) // 镜像的ImageInspect类型转换Json字符串
		// _ = json.Unmarshal([]byte(imgInspectByte), &imgInspectMap) // 镜像的Json类型转换Map
		json.NewDecoder(strings.NewReader(string(imgInspectJson))).Decode(&imgInspectMap)

		// config := imgInspectMap["Config"]
		config := imgInspectMap["Config"].(map[string]interface{})
		if config["Cmd"] != nil {
			// var cmd []string
			// value := reflect.ValueOf(config["Cmd"])
			// for i := 0; i < value.Len(); i++ {
			// 	cmd = append(cmd, value.Index(i).Interface().(string))
			// }
			// fmt.Printf("%#v", cmd)
			item["cmd"] = config["Cmd"]
		}

		item["tty"] = config["Tty"]
		item["open_stdin"] = config["OpenStdin"]
		item["architecture"] = imgInspectMap["Architecture"]
		item["os"] = imgInspectMap["Os"]
		if config["ExposedPorts"] != nil {
			item["ports"] = PortConvert(config["ExposedPorts"].(map[string]interface{}))
		} else {
			item["ports"] = ""
		}
	}

	if 1 == 0 {
		mapJson, _ := json.Marshal(imgMap)
		fmt.Printf("镜像数据的json   %v\n", string(mapJson))

	}
	return item
}
