package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kalougata/bookkeeping/configs"
)

type Config struct {
	DB  *configs.Database
	RDB *configs.Redis
	JWT *configs.JWT
}

func NewConfig() *Config {
	return &Config{
		DB:  configs.DatabaseConfig(),
		RDB: configs.RedisConfig(),
		JWT: configs.JWTConfig(),
	}
}
