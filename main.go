package main

import (
	"log"

	"github.com/bluesky/docker-go-api/container"
	"github.com/bluesky/docker-go-api/image"
	"github.com/docker/docker/client"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

/*
程序不使用root权限，需要将docker的服务端口暴露出来 不过此方法不安全，不需要通过验证就可以直接调用docker了
该程序仅供当前用户使用，所以将当前用户加入到docker用户组中，不需要使用root权限也可以调用docker。

暴露端口
str="/usr/bin/dockerd -H tcp://localhost:2375 -H unix:///var/run/docker.sock "
sudo sed -i "s@/usr/bin/dockerd@$str@" /usr/lib/systemd/system/docker.service
sudo systemctl daemon-reload && systemctl restart docker

将登录用户加入到docker用户组中
sudo gpasswd -a $USER docker && newgrp docker
*/

var (
	// 传入环境变量，以及版本号，初始化一个新的API客户端。如果版本号为空，它不会发送任何版本信息。
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
)

func main() {
	service, err := dbusutil.NewSessionService()
	if err != nil {
		log.Panic("dbus服务初始化失败")
	}

	con := container.NewContainerService(service, cli)
	log.Println("容器服务启动成功")

	img := image.NewImageService(service, cli)
	log.Println("镜像服务启动成功")

	service.Wait()
}
