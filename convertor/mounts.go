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
		// fmt.Printf("%#v\n", mount)
		item := map[string]interface{}{
			"rw":   mount["RW"],
			"dest": mount["Destination"],
			"mode": mount["Mode"],
			"src":  mount["Source"],
			"prop": mount["Propagation"],
			"type": mount["Type"],
		}
		// 这两个字段存在空值情况单独挑出来判断
		if mount["Name"] == nil {
			item["name"] = ""
		} else {
			item["name"] = mount["Name"]
		}

		if mount["driver"] == nil {
			item["driver"] = ""
		} else {
			item["driver"] = mount["Driver"]
		}
		result = append(result, item)
	}
	return result
}
