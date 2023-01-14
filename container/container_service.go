package container

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
	"golang.org/x/vuln/client"
)

const (
	dbusPath        = "/com/bluesky/docker/Container"
	dbusServiceName = "com.bluesky.docker.Container"
	dbusInterface   = dbusServiceName
)

type Container struct {
	service *dbusutil.Service
	cli     *client.Client
}
