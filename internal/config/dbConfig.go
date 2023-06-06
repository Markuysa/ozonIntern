package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}

// New creates new db config
func New() (*Config, error) {
	filePath := "config.yml"
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("cannot find config file")
	}
	return &Config{
		DBName:   viper.GetString("database.name"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
	}, nil
}
