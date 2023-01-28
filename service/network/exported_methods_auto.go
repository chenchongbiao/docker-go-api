package network

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (net *NetworkService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetNetworkList", // 获取镜像列表
			Fn:      net.GetNetworkList,
			OutArgs: []string{"result"},
		},
	}
}
