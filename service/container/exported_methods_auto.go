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
		{
			Name:   "StartContainer",
			Fn:     con.StartContainer,
			InArgs: []string{"ids"},
			// OutArgs: []string{"result"},
		},
		{
			Name:   "StopContainer",
			Fn:     con.StopContainer,
			InArgs: []string{"ids"},
			// OutArgs: []string{"result"},
		},
		{
			Name:   "RestartContainer",
			Fn:     con.RestartContainer,
			InArgs: []string{"ids"},
			// OutArgs: []string{"result"},
		},
	}
}
