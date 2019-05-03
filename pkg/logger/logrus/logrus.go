package logrus

import (
	"github.com/panace9i/galileo/pkg/logger"
	"github.com/sirupsen/logrus"
)

func New(config *logger.Config) *logrus.Logger {
	logger := logrus.New()
	logger.Level = logrusLevelConverter(config.Level)
	logger.WithFields(logrus.Fields(config.Fields))
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: !config.Time,
		TimestampFormat:  config.TimeFormat,
	})

	return logger
}

func logrusLevelConverter(level logger.Level) logrus.Level {
	switch level {
	case logger.LevelDebug:
		return logrus.DebugLevel
	case logger.LevelInfo:
		return logrus.InfoLevel
	case logger.LevelWarn:
		return logrus.WarnLevel
	case logger.LevelError:
		return logrus.ErrorLevel
	case logger.LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
