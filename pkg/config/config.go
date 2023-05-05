package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server struct {
			APP_NAME           string `mapstructure:"name"`
			VERSION            string `mapstructure:"version"`
			HOST               string `mapstructure:"host"`
			INTERNAL_HTTP_PORT int    `mapstructure:"port_rest_internal"`
			EXTERNAL_HTTP_PORT int    `mapstructure:"port_rest_external"`
			GRPC_PORT          int    `mapstructure:"port_grpc"`
			LOG_LEVEL          string `mapstructure:"logger_level"`
		} `mapstructure:"server"`
		Profile struct {
			MYSQL_DSN string `mapstructure:"mysql_dsn"`
		} `mapstructure:"profile"`
	}
)

func NewConfig() (*Config, error) {
	var config Config
	workdir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Search config in home directory with name ".server" (without extension).
	viper.AddConfigPath(workdir + "/configs")
	viper.SetConfigType("toml")
	// viper.SetConfigName("config.toml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	err = viper.Unmarshal(&config)

	return &config, err
}
