package config

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/spf13/viper"
)

func Read() (todoapp.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return todoapp.Config{}, err
	}

	var config todoapp.Config
	if err := viper.Unmarshal(&config); err != nil {
		return todoapp.Config{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}
