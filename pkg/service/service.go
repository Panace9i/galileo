package service

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/handlers"
	"github.com/panace9i/galileo/pkg/logger"
	"github.com/panace9i/galileo/pkg/logger/logrus"
	"github.com/panace9i/galileo/pkg/storage"
	"github.com/panace9i/galileo/pkg/system"
)

func Setup(cfg *config.Config) (r *fasthttprouter.Router, log logger.Logger, sh system.Operator, err error) {
	log = logrus.New(&logger.Config{
		Level: cfg.LogLevel,
		Time:  cfg.ShowLogTime,
	})

	s := storage.New(cfg)
	h := handlers.New(log, cfg, &s)

	r = fasthttprouter.New()
	r.NotFound = h.Wrap(h.NotFound)
	r.MethodNotAllowed = h.Wrap(h.MethodNotAllowed)
	r.PanicHandler = h.PanicHandler
	r.POST("/users/registration", h.Wrap(h.UsersRegistration))
	r.GET("/users/devices", h.WrapAuth(h.UsersDevices))
	r.POST("/devices/registration", h.Wrap(h.DevicesRegistration))
	r.POST("/devices", h.WrapAuth(h.DevicesStats))
	r.GET("/devices/:id", h.WrapAuth(h.Devices))
	r.GET("/info", h.Info)

	sh = system.New(&s, cfg)

	return
}
