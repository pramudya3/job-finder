package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string `env:"ENV"`
	CtxTimeout  int    `env:"CTX_TIMEOUT,required"`
	ServerAddr  string `env:"SERVER_ADDR,required"`
	DatabaseUri string `env:"DATABASE_URI,required"`
	RedisUri    string `env:"REDIS_URI,required"`
	AuthSecret  string `env:"AUTH_SECRET,required"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
