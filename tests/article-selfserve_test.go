package tests

import (
	"github.com/yangpeng-chn/go-web-framework/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// there is no need to run another http server to run the test cases in this file
// the response body will be shown in the same console as well as the test result

func TestGetArticlesHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles", nil)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder() // return httptest.ResponseRecorder
	handler := http.HandlerFunc(controllers.GetArticlesHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", rr.Code, http.StatusOK)
	}

	expected := `[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]`
	if rr.Body.String() != expected {
		t.Errorf("Success expected: %v returned: %v", expected, rr.Body.String())
	}
}

func TestGetArticleHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles/1", nil)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetArticleHandler)
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, rr.Code)
	}

	expected := `{"id":1,"title":"title1","content":"content1"}`
	if rr.Body.String() != expected {
		t.Errorf("Success expected: %v returned: %v", expected, rr.Body.String())
	}
}

func TestAddArticleHandler(t *testing.T) {
	dataJson := `{
 "id": 4,
 "title": "title4",
 "content": "content4"
 }`
	reader := strings.NewReader(dataJson) //Convert string to reader, return strings.Reader
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "http://localhost:4201/v1/articles", reader)
	handler := http.HandlerFunc(controllers.AddArticleHandler)
	handler.ServeHTTP(rr, request)

	if err != nil {
		t.Error(err)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateArticleHandler(t *testing.T) {
	dataJson := `{
 "id": 4,
 "title": "updated-title",
 "content": "updated-content"
 }`
	reader := strings.NewReader(dataJson)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", "http://localhost:4201/v1/articles/1", reader)
	handler := http.HandlerFunc(controllers.UpdateArticleHandler)
	handler.ServeHTTP(rr, request)

	if err != nil {
		t.Error(err)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, rr.Code)
	}
}

func TestDeleteArticleHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	request, err := http.NewRequest("DELETE", "http://localhost:4201/v1/articles/4", nil)
	handler := http.HandlerFunc(controllers.DeleteArticleHandler)
	handler.ServeHTTP(rr, request)
	if err != nil {
		t.Error(err)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, rr.Code)
	}
}
