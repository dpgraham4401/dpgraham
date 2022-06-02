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
	} else if r.URL.Path == "/about" || r.URL.Path == "/about/" {
		view := views.About
		renderTemplate(w, view)
	} else if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
	} else {
		allArticles, err := loadArticles()
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
		}
		allArticles.Articles = allArticles.Articles[:3]
		for i := range allArticles.Articles {
			err := allArticles.Articles[i].loadContent()
			if err != nil {
				errorHandler(w, http.StatusInternalServerError)
			}
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
	allArticles, err := loadArticles()
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
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
			err := art.loadContent()
			if err != nil {
				errorHandler(w, http.StatusInternalServerError)
			}
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
