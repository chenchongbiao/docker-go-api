package convertor

import "github.com/docker/docker/client"

type VolumeConvertor struct {
	cli *client.Client
}

func NewVolumeConvert(cli *client.Client) *VolumeConvertor {
	volumeConvertor := VolumeConvertor{
		cli: cli,
	}
	return &volumeConvertor
}

func (v *VolumeConvertor) VolumeConvert(volume map[string]interface{}, verbose bool) map[string]interface{} {
	item := map[string]interface{}{
		"name":        volume["Name"],
		"driver":      volume["Driver"],
		"mount_point": volume["Mountpoint"],
		"scope":       volume["Scope"],
		"create_time": volume["CreatedAt"],
	}

	// if verboes == true {
	// 	item["driver_opts"] = volume["Options"]
	// }
	return item
}
