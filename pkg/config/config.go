package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/panace9i/galileo/pkg/logger"
)

const (
	SERVICENAME = "GALILEO"
)

type Config struct {
	AppVersion  string       `split_words:"true"`
	LocalHost   string       `split_words:"true"`
	LocalPort   int          `split_words:"true"`
	LogLevel    logger.Level `split_words:"true"`
	ShowLogTime bool         `split_words:"true"`
	DumpPath    string       `split_words:"true"`
}

func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
