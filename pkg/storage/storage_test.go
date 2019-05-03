package storage

import (
	"github.com/panace9i/galileo/pkg/config"
	"testing"
)

func TestNew(t *testing.T) {
	s := New(new(config.Config))
	if s.users == nil {
		t.Error("Expected new users, got nil")
	}
	if s.devices == nil {
		t.Error("Expected new devices, got nil")
	}
	if s.deviceStats == nil {
		t.Error("Expected new deviceStats, got nil")
	}
	if s.userDevices == nil {
		t.Error("Expected new userDevices, got nil")
	}
}
