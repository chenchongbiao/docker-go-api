package convertor

// type PortsMapConvertor struct {
// 	cli *client.Client
// }

// func NewPortsMapConvertor(cli *client.Client) *PortsMapConvertor {
// 	portsMapConvertor := PortsMapConvertor{
// 		cli: cli,
// 	}
// 	return &portsMapConvertor
// }

// 将端口映射的数据重新构造
func PortsMapConvert(portsMap []interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	for _, value := range portsMap {
		mapp := value.(map[string]interface{})
		result = append(result, map[string]interface{}{
			"port":        mapp["PrivatePort"],
			"protocol":    mapp["Type"],
			"listen_ip":   mapp["IP"],
			"listen_port": mapp["PublicPort"],
		})
	}
	return result
}
