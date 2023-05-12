package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		APP_NAME           string `mapstructure:"SERVER_NAME"`
		VERSION            string `mapstructure:"VERSION"`
		HOST               string `mapstructure:"HOST"`
		INTERNAL_HTTP_PORT int    `mapstructure:"PORT_REST_EXTERNAL"`
		EXTERNAL_HTTP_PORT int    `mapstructure:"PORT_REST_INTERNAL"`
		GRPC_PORT          int    `mapstructure:"PORT_GRPC"`
		LOG_LEVEL          string `mapstructure:"LOGGER_LEVEL"`
		PROFILE_MYSQL_DSN  string `mapstructure:"PROFILE_MYSQL_DSN"`
		TASKS_MYSQL_DSN    string `mapstructure:"TASKS_MYSQL_DSN"`
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
	viper.SetConfigType("env")
	// viper.SetConfigName("config.toml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	err = viper.Unmarshal(&config)

	return &config, err
}
