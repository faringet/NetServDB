package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string  `mapstructure:"PORT"`
	Redis   Redis   `mapstructure:"REDIS"`
	Postgre Postgre `mapstructure:"POSTGRE"`
	Auth    Auth    `mapstructure:"AUTH"`
}

type Redis struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Postgre struct {
	DbURL string `mapstructure:"DB_URL"`
}

type Auth struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
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
