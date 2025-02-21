package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shegai01/Example-RestAPI/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	// на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "./config/api.toml", "path to config file in .toml format")
}
func main() {
	//инициализации переменной configPath значением
	flag.Parse()
	log.Println("It works")

	//server instance initialzation
	config := api.NewConfig()
	//теперь тут надо попробовать прочитать из .toml/.env
	_, err := toml.DecodeFile(configPath, config) // дессериализуем содержимое .toml
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
