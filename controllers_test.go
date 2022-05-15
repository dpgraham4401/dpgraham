package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func setUp() (*http.ServeMux, *httptest.ResponseRecorder) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog/", blogHandler)
	writer = httptest.NewRecorder()
	return mux, writer
}

func TestHomeHandlerReturns200(t *testing.T) {
	mux, writer = setUp()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("got non-2XX status code")
	}
}

func TestBlog(t *testing.T) {
	mux, writer = setUp()
	request, _ := http.NewRequest("GET", "/blog/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("got non-2XX status code")
	}
}

func TestPOSTToBlogReturns405(t *testing.T) {
	mux, writer = setUp()
	request, err := http.NewRequest("POST", "/blog/first_post", nil)
	if err != nil {
		t.Fatal(err)
	}
	mux.ServeHTTP(writer, request)
	if writer.Code != 405 {
		t.Errorf("expected %d, recieved %d", http.StatusMethodNotAllowed, writer.Code)
	}
}

func TestPOSTToHomeReturns405(t *testing.T) {
	mux, writer = setUp()
	request, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	mux.ServeHTTP(writer, request)
	if writer.Code != 405 {
		t.Errorf("expected %d, recieved %d", http.StatusMethodNotAllowed, writer.Code)
	}
}

func TestUnknownBlogReturns404(t *testing.T) {
	mux, writer = setUp()
	request, err := http.NewRequest("GET", "/blog/blah", nil)
	if err != nil {
		t.Fatal(err)
	}
	mux.ServeHTTP(writer, request)
	if status := writer.Code; status != http.StatusNotFound {
		t.Errorf("expected %d, recieved %d", status, http.StatusNotFound)
	}
}

func TestUnknownHomeReturns404(t *testing.T) {
	mux, writer = setUp()
	request, err := http.NewRequest("GET", "/blah", nil)
	if err != nil {
		t.Fatal(err)
	}
	mux.ServeHTTP(writer, request)
	if status := writer.Code; status != http.StatusNotFound {
		t.Errorf("expected %d, recieved %d", status, http.StatusNotFound)
	}
}
