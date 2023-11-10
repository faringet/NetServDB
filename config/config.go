package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port      int    `mapstructure:"PORT"`
	DBURL     string `mapstructure:"DB_URL"`
	Username  string `mapstructure:"USERNAME"`
	Password  string `mapstructure:"PASSWORD"`
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Чтение конфигурации из файла
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Раскодирование конфигурации в структуру
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}
