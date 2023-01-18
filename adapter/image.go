package adapter

import (
	"context"
	"strings"

	"github.com/bluesky/docker-go-api/convertor"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ImageAdapter struct {
	cli       *client.Client
	convertor *convertor.ImageConvertor
}

func NewImageAdapter(cli *client.Client) *ImageAdapter {
	imgAdapter := ImageAdapter{
		cli:       cli,
		convertor: convertor.NewImageConvertor(cli),
	}

	return &imgAdapter
}

/*
传入id返回对应的镜像的数据
*/
func (i *ImageAdapter) Item(id string) map[string]interface{} {
	// var imgMap map[string]interface{} // 保存镜像的结构体转换的map

	/*
		获取镜像列表
		尝试加入id字段筛选出镜像，在容器列表尝试可以，
		filter := filters.NewArgs()
		filter.Add("id", id)
		这里并不行，先使用循环查找
	*/
	images, _ := i.cli.ImageList(context.Background(), types.ImageListOptions{})
	/*
		遍历获取的镜像列表传入转换器，获取到map[string]interface{}类型的数据
		取出map数据id对应的值 通过strings.Contains判断获取的id，是否包含传入的id，传入的id一般是id数据的前10位，
		但是不确定是否可能，本地存在，中间部分包含传入的id的镜像
	*/
	for index := range images {
		item := i.convertor.ImageConvert(images[index], false)
		// 找到id对应的镜像
		if strings.Contains(item["id"].(string), id) {
			return item
		}
	}
	if 1 == 0 {
		// imgsJson, _ := json.Marshal(images) // 镜像的ImageSummary类型转换编码成json字符串
		// _ = json.Unmarshal([]byte(imgByte), &imgMap) // 镜像的json字符串类型转换Map
		// json.NewDecoder(strings.NewReader(string(imgsJson))).Decode(&imgMap)

	}
	return nil
}

/*
传入转换器，构造数据
*/
func (i *ImageAdapter) Convert(imageSummary types.ImageSummary, verbose bool) map[string]interface{} {
	return i.convertor.ImageConvert(imageSummary, verbose)
}

func (i *ImageAdapter) List(imagesSummary []types.ImageSummary) []map[string]interface{} {
	items := make([]map[string]interface{}, 0, len(imagesSummary))
	for index := range imagesSummary {
		item := i.convertor.ImageConvert(imagesSummary[index], true)
		// fmt.Printf("%#v\n", item)
		items = append(items, item)
	}
	// fmt.Printf("%#v", items)
	return items
}
