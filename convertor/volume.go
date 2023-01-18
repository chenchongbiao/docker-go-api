package convertor

import (
	"fmt"

	"github.com/docker/docker/client"
)

type VolumeConvertor struct {
	cli *client.Client
}

func NewVolumeConvert(cli *client.Client) *VolumeConvertor {
	volumeConvertor := VolumeConvertor{
		cli: cli,
	}
	return &volumeConvertor
}

func (v *VolumeConvertor) VolumeConvert(volume map[string]interface{}) map[string]interface{} {
	fmt.Printf("%#v", volume)
	return nil
}
