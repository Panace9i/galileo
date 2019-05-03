package storage

import (
	"crypto/md5"
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"testing"
	"time"
)

func TestNewDevice(t *testing.T) {
	email := "123@gmail.com"
	name := "test"
	hash := md5.Sum([]byte(email))
	ut := fmt.Sprintf("%x", hash)

	hash = md5.Sum([]byte(email + name))
	dt := fmt.Sprintf("%x", hash)

	s := New(new(config.Config))
	result := s.NewDevice(email, name)
	if dt != result {
		t.Errorf("Expected %s, got %s", dt, result)
	}

	if s.userDevices[ut][0] != dt {
		t.Errorf("Expected %s, got %s", dt, s.userDevices[ut][0])
	}
}

func TestAddStat(t *testing.T) {
	var collection []DeviceStat
	now := time.Now().Unix()
	token := "123"
	temp := int32(1231)
	collection = append(collection, DeviceStat{
		now,
		temp,
	})

	s := New(new(config.Config))
	s.AddStat(token, collection)
	if s.deviceStats[token][0].Time != now {
		t.Errorf("Expected %d, got %d", now, s.deviceStats[token][0].Time)
	}
	if s.deviceStats[token][0].Temp != temp {
		t.Errorf("Expected %d, got %d", temp, s.deviceStats[token][0].Temp)
	}
}

func TestGetDeviceStat(t *testing.T) {
	var collection []DeviceStat
	now := time.Now().Unix()
	dt := "123"
	ut := "321"
	temp := int32(1231)
	collection = append(collection, DeviceStat{
		now,
		temp,
	})

	s := New(new(config.Config))
	s.AddStat(dt, collection)
	s.userDevices[ut] = append(s.userDevices[ut], dt)

	result := s.GetDeviceStat(ut)
	if result[0].Time != now {
		t.Errorf("Expected %d, got %d", now, result[0].Time)
	}
	if result[0].Temp != temp {
		t.Errorf("Expected %d, got %d", temp, result[0].Temp)
	}
}

func TestDeviceExists(t *testing.T) {
	s := New(new(config.Config))
	result := s.NewDevice("123", "test")

	if !s.DeviceExists(result) {
		t.Errorf("Expected %s, got %s", "'device exists'", "'device not found'")
	}
}
