package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/shegai01/Example-RestAPI/internal/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "config/api.toml", "not finded file .toml format")
}
func main() {
	fmt.Println("hello this is api")
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("not value")
		return
	}
	server := api.New(config)
	if err := server.Start(); err != nil {
		log.Printf("start api failed %v:", err)
		return
	}
	log.Println("connection OK!")
}
