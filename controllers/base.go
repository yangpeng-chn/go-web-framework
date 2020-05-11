package controllers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite database driver
	"github.com/yangpeng-chn/go-web-framework/models"
	"github.com/yangpeng-chn/go-web-framework/utils/logger"
	"github.com/yangpeng-chn/go-web-framework/utils/settings"
)

// Server defines server object containing DB and router instances
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// InitializeDB inits db connection
func (server *Server) InitializeDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "sqlite3" {
		//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DbName)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
		server.DB.Exec("PRAGMA foreign_keys = ON")
	}
	server.DB.Debug().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(&models.User{}, &models.Post{}) //database migration
}

// Run starts running the server at given port
func (server *Server) Run(addr string) {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	conf := settings.GetConfig()
	if conf.EnableHTTPS {
		fmt.Printf("%s Enabling HTTPS ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("%s Listening on port 4201 ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		log.Fatal(http.ListenAndServeTLS(addr, conf.Cert, conf.Key, server.Router))
	} else {
		fmt.Printf("%s Listening on port 4201 ... [OK]\n", time.Now().Format("2006-01-02 15:04:05"))
		log.Fatal(http.ListenAndServe(addr, server.Router))
	}

}

// ParseRequest read data from request
func (server *Server) ParseRequest(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	// body, err := ioutil.ReadAll(r.Body)
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

// GetCurrentFuncName get current function name for logging
func (server *Server) GetCurrentFuncName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	dotPos := strings.LastIndex(f.Name(), ".")
	return f.Name()[dotPos+1:]
}

// Preflight set Preflight for CORS
func (server *Server) Preflight(w http.ResponseWriter, r *http.Request) {
	logger.WriteLog(r, http.StatusOK, nil, server.GetCurrentFuncName())
}
