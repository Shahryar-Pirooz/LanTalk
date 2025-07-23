package config

type Config struct {
	NatsServer NatsServer `mapstructure:"nats"`
}

type NatsServer struct {
	Address string `mapstructure:"address"`
	Port    uint   `mapstructure:"port"`
}
