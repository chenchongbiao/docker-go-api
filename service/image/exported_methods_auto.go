package image

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (img *ImageService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetImageList", // 获取镜像列表
			Fn:      img.GetImageList,
			OutArgs: []string{"result"},
		},
		{
			Name:    "SearchImageById", // 根据镜像id获取镜像
			Fn:      img.SearchImageById,
			InArgs:  []string{"id"},
			OutArgs: []string{"result"},
		},
	}
}
