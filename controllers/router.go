package controllers

import (
	"github.com/yangpeng-chn/go-web-framework/middlewares"
)

func (s *Server) initializeRoutes() {
	// Preflight handler for all preflight requests
	s.Router.Methods("OPTIONS").HandlerFunc(middlewares.SetMiddlewareJSON(s.Preflight))

	// Login Route
	// s.Router.HandleFunc("/v1/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Posts routes, in database
	s.Router.HandleFunc("/v1/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/v1/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/v1/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/v1/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/v1/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	// Articles routes local database
	s.Router.HandleFunc("/v1/articles", middlewares.SetMiddlewareJSON(s.AddArticle)).Methods("POST")
	s.Router.HandleFunc("/v1/articles", middlewares.SetMiddlewareJSON(s.GetArticles)).Methods("GET")
	s.Router.HandleFunc("/v1/articles/{id}", middlewares.SetMiddlewareJSON(s.GetArticle)).Methods("GET")
	s.Router.HandleFunc("/v1/articles/{id}", middlewares.SetMiddlewareJSON(s.UpdateArticle)).Methods("PUT")
	s.Router.HandleFunc("/v1/articles/{id}", middlewares.SetMiddlewareJSON(s.DeleteArticle)).Methods("DELETE")
}
