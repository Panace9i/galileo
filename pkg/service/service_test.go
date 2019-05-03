package service

import (
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/handlers"
	"github.com/panace9i/galileo/pkg/storage"
	"testing"
)

func TestSetup(t *testing.T) {
	cfg := new(config.Config)
	err := cfg.Load(config.SERVICENAME)
	if err != nil {
		t.Error("Error while getting config ", err)
	}
	router, logger, sh, err := Setup(cfg)
	if err != nil {
		t.Errorf("Fail, got '%s', want '%v'", err, nil)
	}
	if router == nil {
		t.Error("Expected new router, got nil")
	}
	if logger == nil {
		t.Error("Expected new logger, got nil")
	}
	if sh == nil {
		t.Error("Expected new SignalHandler, got nil")
	}

	s := storage.New(cfg)
	h := handlers.New(logger, cfg, &s)
	if h == nil {
		t.Error("Expected new handler, got nil")
	}
}
