package main

import (
	"bufio"
	"flag"
	"fmt"
	"lantak/config"
	"lantak/pkg/mo"
	"os"
)

func main() {
	cfg := getConfigFile()
	subj, user := promptUserInput()
	m := mo.NewNats(user, subj, cfg)
	m.Sub()
	handleInput(m)
}

func promptUserInput() (string, string) {
	var subj, user string
	fmt.Print("\nWhat is your subject: ")
	fmt.Scan(&subj)
	fmt.Print("\nEnter your name: ")
	fmt.Scan(&user)
	return subj, user
}

func handleInput(m *mo.Server) {
	for {
		buf := bufio.NewReader(os.Stdin)
		i, _ := buf.ReadString('\n')
		if i == "/exit\n" {
			m.Disconnect()
			break
		}
		m.Pub(i)
	}
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
