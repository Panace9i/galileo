package storage

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type deviceInfo struct {
	Id               string
	Name             string
	RegistrationDate int64
	LastStatUpdate   int64
}

type DeviceStat struct {
	Time int64
	Temp int32
}

func (s *Storage) NewDevice(email string, name string) string {
	hash := md5.Sum([]byte(email + name))
	token := fmt.Sprintf("%x", hash)

	if s.DeviceExists(token) {
		return token
	}

	s.Lock()
	s.devices[token] = deviceInfo{
		Id:               token,
		Name:             name,
		RegistrationDate: time.Now().Unix(),
	}
	s.Unlock()

	s.addUserDevice(email, token)

	return token
}

func (s *Storage) GetDeviceStat(token string) (collection []DeviceStat) {
	s.RLock()
	for _, d := range s.userDevices[token] {
		for _, ds := range s.deviceStats[d] {
			collection = append(collection, ds)
		}
	}
	s.RUnlock()
	return
}

func (s *Storage) AddStat(token string, collection []DeviceStat) {
	s.Lock()

	for _, stat := range collection {
		s.deviceStats[token] = append(s.deviceStats[token], stat)
	}

	s.devices[token] = deviceInfo{
		Id:               s.devices[token].Id,
		Name:             s.devices[token].Name,
		RegistrationDate: s.devices[token].RegistrationDate,
		LastStatUpdate:   time.Now().Unix(),
	}

	s.Unlock()
}

func (s *Storage) DeviceExists(token string) bool {
	s.RLock()
	defer s.RUnlock()
	return s.devices[token].Name != ""
}

func (s *Storage) SaveDevices(fileName string) {
	u, err := json.Marshal(s.devices)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}

	err = ioutil.WriteFile(fileName, u, 0644)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}
}

func (s *Storage) SaveDeviceStats(fileName string) {
	u, err := json.Marshal(s.deviceStats)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}

	err = ioutil.WriteFile(fileName, u, 0644)
	if err != nil {
		fmt.Printf("'%s' is not saved by error: %v\n", fileName, err)
		return
	}
}
