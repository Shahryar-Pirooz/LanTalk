package config

type Config struct {
	HttpServer HttpServer `mapstructure:"http"`
}

type HttpServer struct {
	Address string `mapstructure:"address"`
	Port    uint   `mapstructure:"port"`
}
