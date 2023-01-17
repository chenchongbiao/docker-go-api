package convertor

import (
	"strconv"
	"strings"
)

func PortConvert(ports map[string]interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	if ports != nil {
		for value := range ports {
			arr := strings.Split(value, "/") // 以 / 切割成字符串数组
			port, _ := strconv.Atoi(arr[0])  // 获取端口转换为整型
			protocol := arr[1]               // 获取协议
			// 将端口和协议转换成map加入到数组中
			result = append(result, map[string]interface{}{
				"port":     port,
				"protocol": protocol,
			})
		}
	}
	return result
}
