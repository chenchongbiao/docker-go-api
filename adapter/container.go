package adapter

import (
	"context"
	"time"

	"github.com/bluesky/docker-go-api/convertor"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type ContainerAdapter struct {
	cli       *client.Client
	convertor *convertor.ContainerConvertor
}

func NewContainerAdapter(cli *client.Client) *ContainerAdapter {
	conAdapter := ContainerAdapter{
		cli:       cli,
		convertor: convertor.NewContainerConvertor(cli),
	}
	return &conAdapter
}

func (c *ContainerAdapter) Item(id string) map[string]interface{} {
	filter := filters.NewArgs()
	filter.Add("id", id)
	containers, _ := c.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true, Filters: filter})

	return c.convertor.ContainerConvert(containers[0], true)
}

func (c *ContainerAdapter) Convert(container types.Container) map[string]interface{} {
	return c.convertor.ContainerConvert(container, true)
}

func (c *ContainerAdapter) List(containers []types.Container) []map[string]interface{} {
	items := make([]map[string]interface{}, 0, len(containers))
	for _, container := range containers {
		item := c.convertor.ContainerConvert(container, true)
		// fmt.Printf("%#v\n", item)
		items = append(items, item)
	}

	return items
}

func (c *ContainerAdapter) Start(id string) bool {
	err := c.cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		return false
	}
	return true
}

func (c *ContainerAdapter) Stop(id string) bool {
	timeout := int(time.Minute * 2)
	err := c.cli.ContainerStop(context.Background(), id, container.StopOptions{Timeout: &timeout})
	if err != nil {
		return false
	}
	return true
}

func (c *ContainerAdapter) Restart(id string) bool {
	timeout := int(time.Minute * 2)

	err := c.cli.ContainerStop(context.Background(), id, container.StopOptions{Timeout: &timeout})
	err = c.cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		return false
	}
	return true
}
