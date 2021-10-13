package devicestore

import (
	"fmt"
	"sync"
)

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

type DeviceStore struct {
	sync.Mutex // not sure what is this

	devices map[int]Device
	nextId  int
}

func New() *DeviceStore {
	devStore := &DeviceStore{}
	devStore.devices = make(map[int]Device)
	devStore.nextId = 0
	return devStore
}

func (ds *DeviceStore) CreateDevice(name, devType, owner, mac, ip, startDate string, isCommonlyUsed bool) Device {
	ds.Lock()
	defer ds.Unlock()

	device := Device{
		Id:             ds.nextId,
		Name:           name,
		DeviceType:     devType,
		Owner:          owner,
		MacAddr:        mac,
		IpAddr:         ip,
		StartUseDate:   startDate,
		IsCommonlyUsed: isCommonlyUsed}

	ds.devices[ds.nextId] = device
	ds.nextId++
	return device
}

func (ds *DeviceStore) DeleteDevice(id int) error {
	ds.Lock()
	defer ds.Unlock()

	if _, ok := ds.devices[id]; !ok {
		return fmt.Errorf("device with id=%d not found", id)
	}

	delete(ds.devices, id)
	return nil
}

func (ds *DeviceStore) GetDevices() []Device {
	ds.Lock()
	defer ds.Unlock()

	devices := make([]Device, 0, len(ds.devices))
	for _, device := range ds.devices {
		devices = append(devices, device)
	}
	return devices
}

func (ds *DeviceStore) UpdateDevice(device Device) error {
	ds.Lock()
	defer ds.Unlock()

	if _, ok := ds.devices[device.Id]; !ok {
		return fmt.Errorf("device with id=%d not found", device.Id)
	}

	ds.devices[device.Id] = device
	return nil
}
