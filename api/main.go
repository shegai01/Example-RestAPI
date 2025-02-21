package main

import (
	"Example/RestAPI/internal/app/api"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string = "config/api.toml"
)

func init() {

}
func main() {
	log.Println("It works")

	//server instance initialzation
	config := api.NewConfig()
	//теперь тут надо попробовать прочитать из .toml/.env
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		fmt.Println("can't find config file... using default values", err)
		log.Fatal(err)
	}
	//  _, err = godotenv.Read(configPath)

	// загрузить github.com/BurnSushi/toml

	server := api.NewAPI(config)

	//api server start

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
