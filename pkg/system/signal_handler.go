package system

import (
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/storage"
)

func New(s *storage.Storage, cfg *config.Config) SignalHandler {
	return SignalHandler{
		storage: s,
		cfg:     cfg,
	}
}

type SignalHandler struct {
	storage *storage.Storage
	cfg     *config.Config
}

func (h SignalHandler) Shutdown() error {
	h.storage.SaveUsers(fmt.Sprintf("%s/%s", h.cfg.DumpPath, "users"))
	h.storage.SaveDevices(fmt.Sprintf("%s/%s", h.cfg.DumpPath, "devices"))
	h.storage.SaveUserDevices(fmt.Sprintf("%s/%s", h.cfg.DumpPath, "userDevices"))
	h.storage.SaveDeviceStats(fmt.Sprintf("%s/%s", h.cfg.DumpPath, "deviceStats"))

	return fmt.Errorf("Shutdown")
}
