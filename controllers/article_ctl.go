package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/yangpeng-chn/go-web-framework/middlewares"
	"github.com/yangpeng-chn/go-web-framework/models"
	"github.com/yangpeng-chn/go-web-framework/utils/formaterror"
	"github.com/yangpeng-chn/go-web-framework/utils/logger"
)

// AddArticle adds an article into the slice
func (server *Server) AddArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte
	var responseCode = http.StatusBadRequest
	var article models.Article
	var formattedError error

	if body, err = server.ParseRequest(w, r); err != nil {
		responseCode = http.StatusUnprocessableEntity
		goto Error
	}
	if err = json.Unmarshal(body, &article); err != nil {
		responseCode = http.StatusUnprocessableEntity
		goto Error
	}
	if err = article.AddArticle(); err != nil {
		formattedError = formaterror.FormatError(err.Error())
		middlewares.ERROR(w, http.StatusInternalServerError, formattedError)
		logger.WriteLog(r, http.StatusInternalServerError, formattedError, server.GetCurrentFuncName())
		return
	}
	responseCode = http.StatusOK
	// WriteResponse(w, responseCode, nil)
	middlewares.JSON(w, http.StatusOK, article)
	logger.WriteLog(r, responseCode, nil, server.GetCurrentFuncName())
	return
Error:
	// WriteErrorResponse(w, responseCode, err)
	middlewares.ERROR(w, responseCode, err)
	logger.WriteLog(r, responseCode, err, server.GetCurrentFuncName())
}

// GetArticles returns all the articles
func (server *Server) GetArticles(w http.ResponseWriter, r *http.Request) { //get all articles
	var formattedError error
	article := models.Article{}
	articles, err := article.GetArticles()
	if err != nil {
		formattedError = formaterror.FormatError(err.Error())
		middlewares.ERROR(w, http.StatusInternalServerError, formattedError)
		logger.WriteLog(r, http.StatusInternalServerError, formattedError, server.GetCurrentFuncName())
		return
	}
	middlewares.JSON(w, http.StatusOK, articles)
	logger.WriteLog(r, http.StatusOK, nil, server.GetCurrentFuncName())
	return
}

// GetArticle return an article by ID
func (server *Server) GetArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	var responseCode = http.StatusBadRequest
	var article *models.Article

	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg, /articles/1, id=1
	if err != nil {
		goto Error
	}

	if article, err = article.GetArticle(id); err != nil {
		if err.Error() == "record not found" {
			responseCode = http.StatusNotFound
		}
		goto Error
	}

	responseCode = http.StatusOK
	middlewares.JSON(w, http.StatusOK, article)
	logger.WriteLog(r, responseCode, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, responseCode, err)
	logger.WriteLog(r, responseCode, err, server.GetCurrentFuncName())
}

// UpdateArticle updates an article by ID
func (server *Server) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	var responseCode = http.StatusBadRequest
	var article models.Article
	var body []byte

	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg
	if err != nil {
		goto Error
	}

	if body, err = server.ParseRequest(w, r); err != nil {
		goto Error
	}
	if err = json.Unmarshal(body, &article); err != nil {
		goto Error
	}

	if err = article.UpdateArticle(id); err != nil {
		goto Error
	}

	responseCode = http.StatusOK
	middlewares.JSON(w, http.StatusOK, article)
	logger.WriteLog(r, responseCode, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, responseCode, err)
	logger.WriteLog(r, responseCode, err, server.GetCurrentFuncName())
}

// DeleteArticle delete an article by ID
func (server *Server) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	var err error
	var responseCode = http.StatusBadRequest
	var article models.Article

	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg
	if err != nil {
		goto Error
	}

	if err = article.DeleteArticle(id); err != nil {
		goto Error
	}

	responseCode = http.StatusOK
	middlewares.JSON(w, responseCode, middlewares.Result{Code: 200, Msg: "OK"})
	logger.WriteLog(r, responseCode, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, responseCode, err)
	logger.WriteLog(r, responseCode, err, server.GetCurrentFuncName())
}
