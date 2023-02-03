package volume

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (vol *VolumeService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetVolumeList", // 获取存储卷列表
			Fn:      vol.GetVolumeList,
			OutArgs: []string{"result"},
		},
	}
}
