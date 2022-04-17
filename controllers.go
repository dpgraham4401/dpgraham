package main

import (
	"github.com/dpgrahm4401/dpgraham.com/views"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	allArticles := loadArticles()
	allArticles.Articles = allArticles.Articles[:3]
	for i, _ := range allArticles.Articles {
		allArticles.Articles[i].loadContent()
	}
	view := views.Index
	allArticles.renderTemplate(w, view)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
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
		art.loadContent()
		art.Content.renderTemplate(w, views.Blog)
	}
}
