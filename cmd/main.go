package main

import (
	"flag"
	"fmt"
	"lantak/config"
	"os"
)

func main() {
	cfg := getConfigFile()
	fmt.Println(cfg)
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
