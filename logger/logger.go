package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/yangpeng-chn/go-web-framework/settings"
)

type Log struct {
	Time         string
	Message      string
	ResponseCode int
	Action       string
	Method       string
	URI          string
	RequestData  string
}

func maskPassword(str string) string {
	re := regexp.MustCompile("\"password\":\\s*\"([^\"]+)\"")
	return re.ReplaceAllString(str, "\"password\": \"*\"")
}

func WriteLog(r *http.Request, code int, erro error, action string) {
	// if r.Method == "OPTIONS" {
	// 	return
	// }

	msg := "OK"
	if erro != nil {
		msg = erro.Error()
	}

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	bodyString := string(bodyBytes)
	bodyString = maskPassword(bodyString)

	logs := Log{
		time.Now().Format("2006-01-02 15:04:05"), msg, code, action, r.Method, r.RequestURI, bodyString,
	}

	b, err := json.Marshal(logs)
	if err != nil {
		log.Fatal(err)
	}

	conf := settings.GetConfig()
	if conf.LogIndent {
		var out bytes.Buffer
		json.Indent(&out, b, "", " ")
		fmt.Printf("%s\n", out.String())
	} else {
		fmt.Printf("%s\n", string(b))
	}
}

func WriteInfoLog(v interface{}) {
	log.Printf("[Info] %v\n", v)
}

func WriteDebugLog(v interface{}) {
	log.Printf("[Debug] %v\n", v)
}
