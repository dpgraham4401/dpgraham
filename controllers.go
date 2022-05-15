package main

import (
	"github.com/dpgrahm4401/dpgraham/views"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"net/http"
)

type errorMsg struct {
	Code    int
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed)
	} else if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
	} else {
		allArticles := loadArticles()
		allArticles.Articles = allArticles.Articles[:3]
		for i := range allArticles.Articles {
			allArticles.Articles[i].loadContent()
		}
		p := bluemonday.StripTagsPolicy()
		for i := range allArticles.Articles {
			stringTest := p.Sanitize(string(allArticles.Articles[i].Content.BodyHTML))
			allArticles.Articles[i].Content.BodyHTML = template.HTML(stringTest)
		}
		view := views.Index
		allArticles.renderTemplate(w, view)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed)
	}
	allArticles := loadArticles()
	var art Article
	if r.URL.Path == "/blog/" {
		allArticles.renderTemplate(w, views.Blogs)
	} else {
		url := r.URL.Path[len("/blog/"):]
		for i, article := range allArticles.Articles {
			if article.URL == url {
				art = allArticles.Articles[i]
			}
		}
		if !art.Publish {
			errorHandler(w, http.StatusNotFound)
		} else {
			art.loadContent()
			art.renderTemplate(w, views.Blog)
		}
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	errorToSend := errorMsg{Code: status, Message: "Unknown Error, Sorry :/"}
	w.WriteHeader(status)
	tmpl := views.Error
	if status == http.StatusNotFound {
		errorToSend.Message = "Resource not found"
		err := tmpl.ExecuteTemplate(w, "base", errorToSend)
		if err != nil {
			http.Error(w, "Error", 500)
		}
	} else if status == http.StatusMethodNotAllowed {
		errorToSend.Message = "Method not allowed"
		err := tmpl.ExecuteTemplate(w, "base", errorToSend)
		if err != nil {
			http.Error(w, "Error", 500)
		}
	} else {
		err := tmpl.ExecuteTemplate(w, "base", errorToSend)
		if err != nil {
			http.Error(w, "Error", 500)
		}

	}
}
