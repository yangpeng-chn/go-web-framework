package tests

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

// to run the test cases in this file, make sure the http server is running
// the response body will be shown in another console other than the one running go test

func TestGetArticlesHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles", nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	// check code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, res.StatusCode) //this means our test failed
	}

	// check body
	expected := `[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]`
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body) //convert io.ReaderCloser to string
	str := buf.String()
	if str != expected {
		t.Errorf("Success expected: %v returned: %v", expected, str) //this means our test failed
	}
}

func TestGetArticleHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:4201/v1/articles/1", nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, res.StatusCode)
	}

	expected := `{"id":1,"title":"title1","content":"content1"}`
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	str := buf.String()
	if str != expected {
		t.Errorf("Success expected: %v returned: %v", expected, str) //this means our test failed
	}
}

func TestAddArticleHandler(t *testing.T) {
	dataJSON := `{
 "id": 4,
 "title": "title4",
 "content": "content4"
 }`
	reader := strings.NewReader(dataJSON) //Convert string to reader
	request, err := http.NewRequest("POST", "http://localhost:4201/v1/articles", reader)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, res.StatusCode)
	}
}

func TestUpdateArticleHandler(t *testing.T) {
	dataJSON := `{
 "id": 4,
 "title": "updated-title",
 "content": "updated-content"
 }`
	reader := strings.NewReader(dataJSON) //Convert string to reader
	request, err := http.NewRequest("PUT", "http://localhost:4201/v1/articles/1", reader)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, res.StatusCode)
	}
}

func TestDeleteArticleHandler(t *testing.T) {
	request, err := http.NewRequest("DELETE", "http://localhost:4201/v1/articles/4", nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d, returned: %d", http.StatusOK, res.StatusCode)
	}
}
