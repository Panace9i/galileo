package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config := new(Config)
	err := config.Load(SERVICENAME)
	if err != nil {
		t.Error("Error while loading config ", err)
	}
}
