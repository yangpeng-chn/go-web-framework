package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yangpeng-chn/go-web-framework/controllers"
	"github.com/yangpeng-chn/go-web-framework/middlewares"
	"gopkg.in/go-playground/assert.v1"
)

var s = controllers.Server{}

// there is no need to run another http server to run the test cases in this file
// the response body will be shown in the same console as well as the test result

func TestGetArticlesServeHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles", nil)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder() // return httptest.ResponseRecorder
	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(s.GetArticles))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]`
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}

func TestGetArticleServeHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles/1", nil)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(s.GetArticle))
	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `{"id":1,"title":"title1","content":"content1"}`
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}

func TestAddArticleServeHandler(t *testing.T) {
	dataJSON := `{
 "id": 4,
 "title": "title4",
 "content": "content4"
 }`
	reader := strings.NewReader(dataJSON) //Convert string to reader, return strings.Reader
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "http://localhost:4201/v1/articles", reader)
	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(s.AddArticle))
	handler.ServeHTTP(rr, request)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateArticleServeHandler(t *testing.T) {
	dataJSON := `{
 "id": 4,
 "title": "updated-title",
 "content": "updated-content"
 }`
	reader := strings.NewReader(dataJSON)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", "http://localhost:4201/v1/articles/1", reader)
	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(s.UpdateArticle))
	handler.ServeHTTP(rr, request)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteArticleServeHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("DELETE", "http://localhost:4201/v1/articles/4", nil)
	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(s.DeleteArticle))
	handler.ServeHTTP(rr, request)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rr.Code)
}
