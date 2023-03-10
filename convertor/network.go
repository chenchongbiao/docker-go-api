package convertor

import (
	"encoding/json"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type NetworkConvertor struct {
	cli *client.Client
}

func NewNetworkConvertor(cli *client.Client) *NetworkConvertor {
	netConveror := NetworkConvertor{
		cli: cli,
	}
	return &netConveror
}

func (n *NetworkConvertor) NetworkConvert(network types.NetworkResource, verbose bool) map[string]interface{} {
	var netMap map[string]interface{} // 存放容器信息的map

	networkJson, _ := json.Marshal(network) // 将容器的结构体数据转换为json数据
	json.NewDecoder(strings.NewReader(string(networkJson))).Decode(&netMap)

	item := map[string]interface{}{
		"id":          netMap["Id"],
		"name":        netMap["Name"],
		"driver":      netMap["Driver"],
		"scope":       netMap["Scope"],
		"create_time": netMap["Created"],
	}
	if len(netMap["Containers"].(map[string]interface{})) == 0 {
		item["container_num"] = 0
	}

	if verbose == true {
		ipam := netMap["IPAM"].(map[string]interface{})

		ipam_cfg_list := ipam["Config"].([]interface{})
		// fmt.Printf("%#v\n", len(ipam_cfg_list))
		if len(ipam_cfg_list) != 0 {
			ipam_cfg := ipam_cfg_list[0].(map[string]interface{})
			item["subnet"] = ipam_cfg["Subnet"]
			item["gateway"] = ipam_cfg["Gateway"]
			item["ip_range"] = ipam_cfg["IPRange"]
		}

		item["ipam_driver"] = ipam["Driver"]
		item["internal"] = ipam["Internal"]
		item["attachable"] = ipam["Attachable"]
		item["options"] = netMap["Options"]
		// item["containers"] = netMap["Containers"] // 没有返回信息暂时不使用
		// containers=[ContainerConvertor.from_docker(i, True) for i in containers]
	}

	// fmt.Printf("%#v\n", item)
	return item
}
