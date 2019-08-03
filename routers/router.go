package routers

import (
	"github.com/gorilla/mux"
	"github.com/yangpeng-chn/go-web-framework/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	//article
	{
		"Add Articles", "POST", "/v1/articles", controllers.AddArticleHandler,
	},
	{
		"Get Articles", "GET", "/v1/articles", controllers.GetArticlesHandler,
	},
	{
		"Get Article", "GET", "/v1/articles/{id}", controllers.GetArticleHandler,
	},
	{
		"Update Article", "PUT", "/v1/articles/{id}", controllers.UpdateArticleHandler,
	},
	{
		"Delete Article", "DELETE", "/v1/articles/{id}", controllers.DeleteArticleHandler,
	},
	{
		"Preflight Articles", "OPTIONS", "/v1/articles", controllers.PreflightHandler,
	},
	{
		"Preflight Article Parameter", "OPTIONS", "/v1/articles/{id}", controllers.PreflightHandler,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
