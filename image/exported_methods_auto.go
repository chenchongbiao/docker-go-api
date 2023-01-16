package image

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (img *ImageService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetImageList",
			Fn:      img.GetImageList,
			OutArgs: []string{"result"},
		},
	}
}
