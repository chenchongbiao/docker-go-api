package convertor

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ImageConvert(cli *client.Client, imageSummary types.ImageSummary, verbose bool) map[string]interface{} {
	ctx := context.Background()
	var imgMap map[string]interface{}        // 保存镜像的结构体转换的map
	var imgInspectMap map[string]interface{} // 保存镜像

	imgByte, _ := json.Marshal(imageSummary)     // 镜像的ImageSummary类型转换byte数组
	_ = json.Unmarshal([]byte(imgByte), &imgMap) // 镜像的byte数组类型转换Map

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
		imgInspect, _, _ := cli.ImageInspectWithRaw(ctx, id)
		imgInspectByte, _ := json.Marshal(imgInspect)              // 镜像的ImageInspect类型转换byte数组
		_ = json.Unmarshal([]byte(imgInspectByte), &imgInspectMap) // 镜像的byte数组类型转换Map
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
		item["ports"] = PortConvert(config["ExposedPorts"].(map[string]interface{}))

	}

	if 1 == 1 {
		// fmt.Printf("镜像数据的json   %v\n", string(imgByte))

	}
	// fmt.Println(string(imgInspectByte))

	// itemByte, err := json.Marshal(item)

	// fmt.Println(string(itemByte))
	return item
}
