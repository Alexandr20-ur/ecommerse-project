package config

import (
	"ecommerce/pkg/log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DBConfig
	Redis RedisConfig
	JWT   JWTConfig
}

type DBConfig struct {
	URL             string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret          string
	AccessTTL       time.Duration
	RefreshTTL      time.Duration
	RefreshRotation bool // выдавать новый refresh при каждом refresh
	Leeway          time.Duration
}

func (config *Config) LoadConfig(logger *log.Logger) error {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failder reading env file: " + err.Error())
		return err
	}

	config.DB = DBConfig{
		URL:             os.Getenv("DATABASE_DSN"),
		MaxConns:        10, // или по формуле ниже
		MinConns:        2,
		MaxConnIdleTime: 5 * time.Minute,
		MaxConnLifetime: 30 * time.Minute,
	}

	config.JWT = JWTConfig{
		Secret:          os.Getenv("JWT_SECRET"),
		AccessTTL:       time.Hour,
		RefreshTTL:      time.Hour,
		Leeway:          30 * time.Second,
		RefreshRotation: true,
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
