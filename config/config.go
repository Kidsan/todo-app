package config

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/spf13/viper"
)

func ReadTodoCLIConfig() (todoapp.CLIConfig, error) {
	return todoapp.CLIConfig{}, nil
}

func Read() (todoapp.APIConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return todoapp.APIConfig{}, err
	}

	var config todoapp.APIConfig
	if err := viper.Unmarshal(&config); err != nil {
		return todoapp.APIConfig{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}
