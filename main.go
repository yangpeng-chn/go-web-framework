package main

import (
	"fmt"
	"github.com/yangpeng-chn/go-web-framework/routers"
	"github.com/yangpeng-chn/go-web-framework/settings"
	"log"
	"net/http"
	"time"
)

func main() {
	conffile := "settings/conf.json"
	if err := settings.ReadConfigFile(conffile); err != nil {
		log.Fatal(err.Error())
		return
	}

	router := routers.NewRouter()
	conf := settings.GetConfig()

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
