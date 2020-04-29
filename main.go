package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yangpeng-chn/go-web-framework/controllers"
	"github.com/yangpeng-chn/go-web-framework/routers"
	"github.com/yangpeng-chn/go-web-framework/seed"
	"github.com/yangpeng-chn/go-web-framework/settings"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func main() {
	var err error
	conffile := "conf/conf.json"
	if err = settings.ReadConfigFile(conffile); err != nil {
		log.Fatal(err.Error())
		return
	}

	router := routers.NewRouter()
	conf := settings.GetConfig()

	if conf.UseDatabase {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		} else {
			fmt.Println("We are getting the env values")
		}

		server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		seed.Load(server.DB)
	}

	//https://gist.github.com/denji/12b3a568f092ab951456
	if conf.EnableHTTPS {
		fmt.Printf("%s Enabling HTTPS ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("%s Listening on port 4201 ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		log.Fatal(http.ListenAndServeTLS(":4201", conf.Cert, conf.Key, router))
	} else {
		fmt.Printf("%s Listening on port 4201 ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		log.Fatal(http.ListenAndServe(":4201", router))
	}
}
