package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	APP_PREFIX = "APP"
)

type AppConfig struct {
	DB DBConfig
}

type DBConfig struct {
}

func ProvideConfig() AppConfig {
	var app AppConfig
	envconfig.MustProcess(APP_PREFIX, &app)
	envconfig.MustProcess(APP_PREFIX, &app.DB)

	return app
}
