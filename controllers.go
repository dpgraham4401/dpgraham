package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadArticle("home")
	renderTemplate(w, "index", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog/" {
		allArticles := loadArticles()
		for _, i := range allArticles.Title {
			fmt.Println(i)
		}
		allArticles.renderTemplate(w, "blog_home")
		// fmt.Println(allArticles.articless)
		// pBlank := Page{}
		// p := &pBlank
		// p.readArticleList()
		// renderTemplate(w, "blog_home", p)
	} else {
		title := r.URL.Path[len("/blog/"):]
		p, _ := loadArticle(title)
		renderTemplate(w, "article", p)
	}
}
