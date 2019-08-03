package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/yangpeng-chn/go-web-framework/settings"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SetResponseHeaders(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	// w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Language, Content-Type")
}

func WriteErrorResponse(w http.ResponseWriter, errCode int, err error) {
	w.WriteHeader(errCode)
	result := Result{Code: errCode, Msg: err.Error()}
	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}

func WriteResponse(w http.ResponseWriter, code int, v interface{}) error {
	res, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.WriteHeader(code) //header
	if v != nil {
		w.Write(res) //body data
	} else {
		result := Result{Code: code, Msg: "OK"}
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}
	return nil
}

func ParseRequest(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10485760)) //10MB
	if err != nil {
		return []byte(""), err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) //add this line to restore the body which will be used in logger.go again

	err = r.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	return body, nil
}

func GetCurrentFuncName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	dotPos := strings.LastIndex(f.Name(), ".")
	return f.Name()[dotPos+1:]
}

func PreflightHandler(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	settings.WriteLog(r, http.StatusOK, nil, GetCurrentFuncName())
}
