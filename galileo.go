package main

import (
	"fmt"
	"github.com/panace9i/galileo/pkg/config"
	"github.com/panace9i/galileo/pkg/service"
	"github.com/panace9i/galileo/pkg/system"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	cfg := new(config.Config)
	if err := cfg.Load(config.SERVICENAME); err != nil {
		log.Fatal(err)
		return
	}

	router, logger, sh, err := service.Setup(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	logger.Info("Version:", cfg.AppVersion)
	logger.Infof("Service %s listened on %s:%d", config.SERVICENAME, cfg.LocalHost, cfg.LocalPort)

	go fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", cfg.LocalHost, cfg.LocalPort), router.Handler)

	signals := system.NewSignals()
	if err := signals.Wait(logger, sh); err != nil {
		logger.Fatal(err)
	}
}
