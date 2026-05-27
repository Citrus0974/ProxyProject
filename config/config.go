package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    appConfig        `yaml:"app"`
		Log    logConfig        `yaml:"log"`
		Server httpServerConfig `yaml:"server"`
		Proxy  proxyConfig      `yaml:"proxy"`
		mutex  sync.Mutex
	}

	appConfig struct {
		Name               string `yaml:"name" env:"APP_NAME" env-default:"ProxyApp"`
		Version            string `yaml:"version" env:"APP_VERSION" env-default:"0.1"`
		ReloadTimerSeconds int    `yaml:"reload-timer" env:"RELOAD_TIMER"`
	}
	logConfig struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"INFO" env-upd:"true"`
	}
	httpServerConfig struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	}
	proxyConfig struct {
		BaseURL string `yaml:"base-url" env:"PROXY_URL" env-default:"http://127.0.0.1:8081"`
		DefaultAllow bool `yaml:"default-allow" env:"DEFAULT_ALLOW" env-default:"false"`
		AllowFile string `yaml:"allow-file" env:"ALLOW_FILE" env-default:"configs/allow.json"`
		DenyFile string `yaml:"deny-file" env:"DENY_FILE" env-default:"configs/deny.json"`
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

func UpdateConfig(cfg *Config) error {
	cfg.mutex.Lock()
	defer cfg.mutex.Unlock()

	err := cleanenv.UpdateEnv(cfg)
	if err != nil {
		return err
	}
	return nil
}
