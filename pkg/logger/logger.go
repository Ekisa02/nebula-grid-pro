package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yourusername/nebula-grid/internal/config"
)

var Log = logrus.New()

func Init(cfg *config.Config) {
	level, err := logrus.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}
