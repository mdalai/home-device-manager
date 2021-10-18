package models

type Devices struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	DeviceType     string `json:"device_type"`
	Owner          string `json:"owner"`
	MacAddr        string `json:"mac_addr"`
	IpAddr         string `json:"ip_addr"`
	StartUseDate   string `json:"start_use_date"`
	IsCommonlyUsed bool   `json:"is_commonly_used"`
}
