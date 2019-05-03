package storage

import (
	"encoding/json"
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"io/ioutil"
	"sync"
)

type Storage struct {
	sync.RWMutex
	users       map[string]userInfo
	userDevices map[string][]string
	devices     map[string]deviceInfo
	deviceStats map[string][]DeviceStat
}

func New(cfg *config.Config) Storage {
	return Storage{
		users:       usersFromFile(fmt.Sprintf("%s/%s", cfg.DumpPath, "users")),
		userDevices: userDevicesFromFile(fmt.Sprintf("%s/%s", cfg.DumpPath, "userDevices")),
		devices:     devicesFromFile(fmt.Sprintf("%s/%s", cfg.DumpPath, "devices")),
		deviceStats: deviceStatsFromFile(fmt.Sprintf("%s/%s", cfg.DumpPath, "deviceStats")),
	}
}

func usersFromFile(fileName string) map[string]userInfo {
	result := make(map[string]userInfo)
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	users := new(map[string]userInfo)
	err = json.Unmarshal(f, users)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	for i, u := range *users {
		result[i] = u
	}

	return result
}

func userDevicesFromFile(fileName string) map[string][]string {
	result := make(map[string][]string)
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	ud := new(map[string][]string)
	err = json.Unmarshal(f, ud)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	for i, u := range *ud {
		result[i] = u
	}

	return result
}

func devicesFromFile(fileName string) map[string]deviceInfo {
	result := make(map[string]deviceInfo)
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	d := new(map[string]deviceInfo)
	err = json.Unmarshal(f, d)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	for i, u := range *d {
		result[i] = u
	}

	return result
}

func deviceStatsFromFile(fileName string) map[string][]DeviceStat {
	result := make(map[string][]DeviceStat)
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	ds := new(map[string][]DeviceStat)
	err = json.Unmarshal(f, ds)
	if err != nil {
		fmt.Printf("'%s' is not loaded by error: %v\n", fileName, err)
		return result
	}

	for i, u := range *ds {
		result[i] = u
	}

	return result
}
