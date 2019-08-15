package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/yangpeng-chn/go-web-framework/models"
	"github.com/yangpeng-chn/go-web-framework/settings"
)

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var body []byte
	var responseCode = http.StatusBadRequest
	var article models.Article

	// or use decode
	// if err = json.NewDecoder(r.Body).Decode(&article); err != nil {
	// 	goto Error
	// }

	if body, err = ParseRequest(w, r); err != nil {
		goto Error
	}
	if err = json.Unmarshal(body, &article); err != nil {
		goto Error
	}
	if err = models.AddArticle(article); err != nil {
		goto Error
	}
	responseCode = http.StatusOK
	WriteResponse(w, responseCode, nil)
	settings.WriteLog(r, responseCode, nil, GetCurrentFuncName())
	return
Error:
	WriteErrorResponse(w, responseCode, err)
	settings.WriteLog(r, responseCode, err, GetCurrentFuncName())
}

func GetArticlesHandler(w http.ResponseWriter, r *http.Request) { //get all articles
	SetResponseHeaders(w, r)
	var err error
	var responseCode = http.StatusBadRequest
	// var articles []*models.Article
	var articles []models.Article

	if articles, err = models.GetArticles(); err != nil {
		goto Error
	}

	if err = WriteResponse(w, http.StatusOK, articles); err != nil {
		goto Error
	}
	responseCode = http.StatusOK
	settings.WriteLog(r, responseCode, nil, GetCurrentFuncName())
	return
Error:
	WriteErrorResponse(w, responseCode, err)
	settings.WriteLog(r, responseCode, err, GetCurrentFuncName())
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var responseCode = http.StatusBadRequest
	var article *models.Article

	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg, /articles/1, id=1
	if err != nil {
		goto Error
	}

	if article, err = models.GetArticle(id); err != nil {
		if err.Error() == "record not found" {
			responseCode = http.StatusNotFound
		}
		goto Error
	}

	if err = WriteResponse(w, http.StatusOK, article); err != nil {
		goto Error
	}

	responseCode = http.StatusOK
	settings.WriteLog(r, responseCode, nil, GetCurrentFuncName())
	return
Error:
	WriteErrorResponse(w, responseCode, err)
	settings.WriteLog(r, responseCode, err, GetCurrentFuncName())
}

func UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var responseCode = http.StatusBadRequest
	var article models.Article
	var body []byte

	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg
	if err != nil {
		goto Error
	}

	if body, err = ParseRequest(w, r); err != nil {
		goto Error
	}
	if err = json.Unmarshal(body, &article); err != nil {
		goto Error
	}

	if err = models.UpdateArticle(id, article); err != nil {
		goto Error
	}

	responseCode = http.StatusOK
	WriteResponse(w, responseCode, nil)
	settings.WriteLog(r, responseCode, nil, GetCurrentFuncName())
	return
Error:
	WriteErrorResponse(w, responseCode, err)
	settings.WriteLog(r, responseCode, err, GetCurrentFuncName())
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w, r)
	var err error
	var responseCode = http.StatusBadRequest

	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	id, err := strconv.Atoi(path.Base(r.URL.Path)) //to use std pkg
	if err != nil {
		goto Error
	}

	if err = models.DeleteArticle(id); err != nil {
		goto Error
	}

	responseCode = http.StatusOK
	WriteResponse(w, responseCode, nil)
	settings.WriteLog(r, responseCode, nil, GetCurrentFuncName())
	return
Error:
	WriteErrorResponse(w, responseCode, err)
	settings.WriteLog(r, responseCode, err, GetCurrentFuncName())
}
