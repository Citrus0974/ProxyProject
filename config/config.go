package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    appConfig        `yaml:"app"`
		Log    logConfig        `yaml:"log"`
		Server httpServerConfig `yaml:"server"`
		Proxy  proxyConfig      `yaml:"proxy"`
	}

	appConfig struct {
		Name    string `yaml:"name" env:"APP_NAME" env-default:"ProxyApp"`
		Version string `yaml:"version" env:"APP_VERSION" env-default:"0.1"`
	}
	logConfig struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"INFO" env-upd:"true"`
	}
	httpServerConfig struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	}
	proxyConfig struct {
		BaseURL string `yaml:"base-url" env:"PROXY_URL" env-default:"http://127.0.0.1:8081"`
	}
)

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
