package adapter

import (
	"github.com/bluesky/docker-go-api/convertor"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type NetworkAdapter struct {
	cli       *client.Client
	convertor *convertor.NetworkConvertor
}

func NewNetworkAdapter(cli *client.Client) *NetworkAdapter {
	netAdapter := NetworkAdapter{
		cli: cli,
		// convertor: convertor.NewImageConvertor(cli),
	}

	return &netAdapter
}

func (n *NetworkAdapter) List(networks []types.NetworkResource) []map[string]interface{} {
	items := make([]map[string]interface{}, 0, len(networks))
	for _, network := range networks {
		item := n.convertor.NetworkConvert(network, true)
		// fmt.Printf("%#v\n", item)
		items = append(items, item)
	}

	return items
}
