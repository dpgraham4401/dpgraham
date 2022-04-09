package main

import (
	"fmt"
	"net/http"
)

func printOrder(list ArticleList) {
	for _, i := range list.Articles {
		fmt.Printf("%s: %d\n", i.Title, i.Id)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	allArticles := loadArticles()
	allArticles.Articles = allArticles.Articles[:3]
	for i, _ := range allArticles.Articles {
		allArticles.Articles[i].loadContent()
	}
	printOrder(allArticles)
	allArticles.renderTemplate(w, "index")
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
