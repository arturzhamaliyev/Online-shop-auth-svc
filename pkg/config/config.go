package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	// DB   struct {
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	// }/
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	fmt.Println(config)
	return config, nil
}
