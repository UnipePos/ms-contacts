package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DataBasePath string      `envconfig:"DATABASE_PATH" default:"database/database.db"`
	DataBaseMode os.FileMode `envconfig:"DATABASE_MODE" default:"0666"`
}

func Load() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
