package models

type TrafficConfig struct {
	// Device 基准服务实例监听网卡
	Device string `json:"device" bson:"device"`
	// Port 基准服务实例监听端口
	Port uint16 `json:"port" bson:"port"`
	// Addr 被测服务地址
	Addr string `json:"addr" json:"addr"`
}
