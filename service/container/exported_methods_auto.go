package container

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (con *ContainerService) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "GetContainerList",
			Fn:      con.GetContainerList,
			InArgs:  []string{"args"},
			OutArgs: []string{"result"},
		},
		{
			Name:    "Item",
			Fn:      con.Item,
			InArgs:  []string{"id"},
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
		{
			Name:   "RmContainer",
			Fn:     con.RmContainer,
			InArgs: []string{"id"},
		},
	}
}
