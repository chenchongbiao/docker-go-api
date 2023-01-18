package convertor

// type MountsConvertor struct {
// 	cli *client.Client
// }

// func NewMountsConvert(cli *client.Client) *MountsConvertor {
// 	mouConvertor := MountsConvertor{
// 		cli: cli,
// 	}
// 	return &mouConvertor
// }

func MountsConvert(mounts []interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	for _, value := range mounts {
		mount := value.(map[string]interface{})
		result = append(result, map[string]interface{}{
			"rw":   mount["RW"],
			"dest": mount["Destination"],
			"mode": mount["Mode"],
			"src":  mount["Source"],
			"prop": mount["Source"],
			"type": mount["Type"],
		})
	}
	return result
}
