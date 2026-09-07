package config

import (
	"ecommerce/pkg/log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DBConfig
	Redis RedisConfig
}

type DBConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func (config *Config) LoadConfig(logger *log.Logger) error {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failder reading env file: " + err.Error())
		return err
	}

	config.DB = DBConfig{
		URL:          os.Getenv("DATABASE_DSN"),
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	return nil
}

func NewConfig(logger *log.Logger) (*Config, error) {
	cfg := &Config{}
	if err := cfg.LoadConfig(logger); err != nil {
		return nil, err
	}

	return cfg, nil
}
