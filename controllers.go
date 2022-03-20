package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadArticle("home")
	renderTemplate(w, "index", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog/" {
		allArticles := loadArticles()
		allArticles.renderTemplate(w, "blog_home")
	} else {
		url := r.URL.Path[len("/blog/"):]
		content, _ := loadContent(url)
		content.renderTemplate(w, "article")
	}
}
