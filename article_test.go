package main

import (
	"testing"
)

func TestLoadArticles(t *testing.T) {
	articles, err := loadArticles()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	articleDates := make([]string, len(articles.Articles))
	for _, a := range articles.Articles {
		articleDates = append(articleDates, a.Date)
	}
	if len(articleDates) == len(articles.Articles) {
		t.Errorf("Number of Dates not equal to number of artciles")
	}
}
