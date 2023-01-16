package container

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (con *ContainerService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetContainerList",
			Fn:      con.GetContainerList,
			OutArgs: []string{"result"},
		},
	}
}