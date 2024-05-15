package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/slogResponse"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	DB     Database `yaml:"db" env-required:"true"`
	Server Server   `yaml:"server" env-required:"true"`
	Nats   Nats     `yaml:"nats" env-required:"true"`
}

type Database struct {
	DbUser  string `yaml:"db_user" env-required:"true"`
	DbPass  string `yaml:"db_pass" env-required:"true"`
	DbName  string `yaml:"db_name" env-required:"true"`
	SslMode string `yaml:"ssl_mode" env-default:"disable"`
	Port    string `yaml:"db_port" env-default:"5432"`
}

type Server struct {
	Timeout     time.Duration `yaml:"timeout" env-default:"15s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
	Address     string        `yaml:"address" env-required:"true"`
}

type Nats struct {
	Address       string        `yaml:"address"`
	Retry         bool          `yaml:"retry"`
	MaxReconnects int           `yaml:"max_reconnects"`
	ReconnectWait time.Duration `yaml:"reconnect_wait"`
}

func MustLoad() *Config {
	const op = "config.MustLoad"

	path, ok := os.LookupEnv("config_path")
	if !ok || path == "" {
		slog.Error("couldn't find config path:", slogResponse.SlogOp(op))
		os.Exit(404)
	}

	if _, err := os.Stat(path); err != nil {
		slog.Error("couldn't find config path: ", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		os.Exit(404)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		slog.Error("couldn't read config", slogResponse.SlogOp(op), slogResponse.SlogErr(err))
		os.Exit(404)
	}

	return &cfg
}
