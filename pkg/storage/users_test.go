package storage

import (
	"crypto/md5"
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"testing"
)

func TestNewUser(t *testing.T) {
	email := "123@gmail.com"
	hash := md5.Sum([]byte(email))
	token := fmt.Sprintf("%x", hash)

	s := New(new(config.Config))
	result := s.NewUser(email)
	if result != token {
		t.Errorf("Expected %s, got %s", token, result)
	}
}

func TestGetUserDevices(t *testing.T) {
	s := New(new(config.Config))
	ut := "123"
	dt := "321"
	name := "test"
	rd := int64(123)

	s.devices[dt] = deviceInfo{
		Id:               dt,
		Name:             name,
		RegistrationDate: rd,
	}

	s.userDevices[ut] = append(s.userDevices[ut], dt)
	result := s.GetUserDevices(ut)
	if result == nil {
		t.Error("Expected user devices, got nil")
	}
	if result[0].Id != dt {
		t.Errorf("Expected %s, got %s", dt, result[0].Id)
	}
	if result[0].Name != name {
		t.Errorf("Expected %s, got %s", name, result[0].Name)
	}
	if result[0].RegistrationDate != rd {
		t.Errorf("Expected %d, got %d", rd, result[0].RegistrationDate)
	}
}

func TestUserExists(t *testing.T) {
	s := New(new(config.Config))
	token := s.NewUser("123")

	if !s.UserExists(token) {
		t.Errorf("Expected %s, got %s", "'user exists'", "'user not found'")
	}
}

func TestUserExistsByEmail(t *testing.T) {
	s := New(new(config.Config))
	email := "123456"
	s.NewUser(email)

	if !s.UserExistsByEmail(email) {
		t.Errorf("Expected %s, got %s", "'user exists'", "'user not found'")
	}
}

func TestAddUserDevice(t *testing.T) {
	s := New(new(config.Config))
	email := "123456"
	hash := md5.Sum([]byte(email))
	ut := fmt.Sprintf("%x", hash)
	dt := "123"

	s.addUserDevice(email, dt)

	if s.userDevices[ut][0] != dt {
		t.Errorf("Expected %s, got %s", dt, s.userDevices[ut][0])
	}
}
