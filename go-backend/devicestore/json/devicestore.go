package jsonstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	models "mdalai/mydeviceservice/model"
)

type DeviceStore struct {
	sync.Mutex // not sure what is this

	jsonFilepath string
}

func New() *DeviceStore {
	devStore := &DeviceStore{jsonFilepath: "db/db.json"}
	return devStore
}

func (ds *DeviceStore) readJsonFile() []models.Device {
	data, err := ioutil.ReadFile(ds.jsonFilepath)
	if err != nil {
		fmt.Print(err)
	}

	var devices models.Devices
	err = json.Unmarshal(data, &devices)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return devices.Devices
}

func (ds *DeviceStore) saveJsonFile(deviceList []models.Device) {
	devices := models.Devices{Devices: deviceList}
	//devicesJson, _ := json.Marshal(devices)
	devicesJson, _ := json.MarshalIndent(devices, "", "    ")
	err := ioutil.WriteFile(ds.jsonFilepath, devicesJson, 0644)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func (ds *DeviceStore) CreateDevice(name, devType, owner, mac, ip, startDate string, isCommonlyUsed bool) models.Device {
	ds.Lock()
	defer ds.Unlock()

	devices := ds.readJsonFile()

	device := models.Device{
		Id:             devices[len(devices)-1].Id + 1,
		Name:           name,
		DeviceType:     devType,
		Owner:          owner,
		MacAddr:        mac,
		IpAddr:         ip,
		StartUseDate:   startDate,
		IsCommonlyUsed: isCommonlyUsed}

	newDevices := append(devices, device)
	ds.saveJsonFile(newDevices)

	return device
}

func (ds *DeviceStore) DeleteDevice(id int) error {
	ds.Lock()
	defer ds.Unlock()

	devices := ds.readJsonFile()
	rmIdx := -1
	for i, device := range devices {
		if device.Id == id {
			rmIdx = i
			break
		}
	}

	var newDevices []models.Device
	// if not found ID, return the same devices
	if rmIdx == -1 {
		newDevices = devices
	} else if rmIdx+1 == len(devices) {
		newDevices = devices[:rmIdx]
	} else if rmIdx == 0 {
		newDevices = devices[rmIdx+1:]
	} else {
		newDevices = append(devices[:rmIdx], devices[rmIdx+1:]...)
	}

	ds.saveJsonFile(newDevices)

	return nil
}

func (ds *DeviceStore) GetDevices() []models.Device {
	ds.Lock()
	defer ds.Unlock()

	devices := ds.readJsonFile()
	return devices
}

func (ds *DeviceStore) UpdateDevice(device models.Device) error {
	ds.Lock()
	defer ds.Unlock()

	devices := ds.readJsonFile()
	var replaceIdx int
	for i, dev := range devices {
		if dev.Id == device.Id {
			replaceIdx = i
			break
		}
	}
	devices[replaceIdx] = device
	ds.saveJsonFile(devices)

	return nil
}
