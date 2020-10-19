package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/apiserver"
)

func main() {
	var configPath string
	loadConfigPath(&configPath)

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

func loadConfigPath(configPath *string) {
	flag.StringVar(configPath, "config-path", "configs/api-server.toml", "path to config file")
	flag.Parse()
}
