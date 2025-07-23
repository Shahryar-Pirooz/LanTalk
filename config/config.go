package config

type Config struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	Address string `mapstructure:"address"`
	Port    uint   `mapstructure:"port"`
}
