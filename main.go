package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yangpeng-chn/go-web-framework/controllers"
	"github.com/yangpeng-chn/go-web-framework/seed"
	"github.com/yangpeng-chn/go-web-framework/settings"
)

var server = controllers.Server{}

func main() {
	var err error
	file := "conf/conf.json"
	if err = settings.ReadConfigFile(file); err != nil {
		log.Fatal(err.Error())
		return
	}

	conf := settings.GetConfig()
	if conf.UseDatabase {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		} else {
			fmt.Println("We are getting the env values")
		}

		server.InitializeDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		seed.Load(server.DB)
	}
	server.Run(":4201")
}
