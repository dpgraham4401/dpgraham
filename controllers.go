package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadArticle("home")
	renderTemplate(w, "index", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	allArticles := loadArticles()
	var art Article
	if r.URL.Path == "/blog/" {
		allArticles.renderTemplate(w, "blog_home")
	} else {
		url := r.URL.Path[len("/blog/"):]
		for i, article := range allArticles.Articles {
			if article.URL == url {
				art = allArticles.Articles[i]
			}
		}
		art.loadContent()
		art.Content.renderTemplate(w, "article")
	}
}
