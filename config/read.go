package config

import (
	"errors"
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)
}

func Read(path string) (Config, error) {
	var cfg Config
	newPath, err := absPath(path)
	if err != nil {
		return cfg, errors.New("failed to get absolute path: " + err.Error())
	}
	viper.AddConfigPath(newPath)
	viper.AutomaticEnv()
	viper.Debug()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return cfg, errors.New("config file not found")
		} else {
			return cfg, errors.New("Config file was found but another error was produced: " + err.Error())
		}
	}
	return cfg, viper.Unmarshal(&cfg)
}

func MustRead(path string) Config {
	cfg, err := Read(path)
	if err != nil {
		panic(err)
	}
	return cfg
}
