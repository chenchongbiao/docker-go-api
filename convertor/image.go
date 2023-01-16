package convertor

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
)

func ImageConvert(imageSummary types.ImageSummary) string {
	var dataMap map[string]interface{}
	imgByte, err := json.Marshal(imageSummary)
	if err != nil {
		log.Println("镜像的ImageSummary类型转换byte数组失败")
	}

	err = json.Unmarshal([]byte(imgByte), &dataMap)
	if err != nil {
		log.Println("镜像的byte数组类型转换Map失败")
	}
	fmt.Printf("字符串 %v\n", string(imgByte))
	// 构造一个新的map
	item := make(map[string]interface{})

	// 获取镜像id
	id, _ := dataMap["Id"].(string)

	// tags, _ := dataMap["RepoTags"]

	// 切割字符串
	item["Id"] = id[7:]

	itemByte, err := json.Marshal(item)

	fmt.Println(string(itemByte))
	return string(itemByte)
}
