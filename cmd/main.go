package main

import (
	"flag"
	"lantak/config"
	"lantak/wire"
	"os"
)

func main() {
	cfg := getConfigFile()
	c := wire.InitChatController(cfg)
	c.PromptUserInput()
	c.HandleInput()
}

func getConfigFile() config.Config {
	var cfgPath string
	const (
		defaultPath = "./config.yaml"
		usage       = "path to config file"
	)
	flag.StringVar(&cfgPath, "c", defaultPath, usage+" (shorthand)")
	flag.StringVar(&cfgPath, "config", defaultPath, usage)
	flag.Parse()
	if envPath := os.Getenv("LANTALK_CONFIG"); envPath != "" {
		cfgPath = envPath
	}
	return config.MustRead(cfgPath)
}
