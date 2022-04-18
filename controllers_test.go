package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog/", blogHandler)
	writer = httptest.NewRecorder()
}

func TestHomeHandlerGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("got non-2XX status code")
	}
}

func TestBlog(t *testing.T) {
	request, _ := http.NewRequest("GET", "/blog/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("got non-2XX status code")
	}
}

func TestBlogFirstPost(t *testing.T) {
	request, _ := http.NewRequest("GET", "/blog/first_post", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("got non-2XX status code")
	}
}
