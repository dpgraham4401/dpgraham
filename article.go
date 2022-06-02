package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

// ArticleList is loaded at runtime and contains slice of articles
type ArticleList struct {
	Articles []Article
}

// Article captures metadata about a blog post or tutorial etc.
type Article struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	LastUpdate string  `json:"lastUpdate"`
	Date       string  `json:"date"`
	Publish    bool    `json:"publish"`
	Type       string  `json:"type"`
	Content    Content `json:"content"`
	URL        string  `json:"url"`
}

// Content captures information on the storage of an article content
type Content struct {
	Body     []byte
	Path     string `json:"path"`
	Format   string `json:"format"`
	BodyHTML template.HTML
}

// loadArticles loads all metadata about the content but not the content itself
func loadArticles() (ArticleList, error) {
	var allArticles ArticleList
	articlePath := "./blog/entries/"
	articleDir, err := os.Open(articlePath)
	if err != nil {
		return allArticles, err
	}
	files, err := articleDir.ReadDir(0)
	if err != nil {
		return allArticles, err
	}
	for _, v := range files {
		var newArticle Article
		file, _ := ioutil.ReadFile(articlePath + v.Name())
		err = json.Unmarshal(file, &newArticle)
		if err != nil {
			return allArticles, err
		}
		allArticles.Articles = append(allArticles.Articles, newArticle)
	}
	sort.Slice(allArticles.Articles, func(i, j int) bool {
		return allArticles.Articles[i].Id > allArticles.Articles[j].Id
	})
	return allArticles, nil
}

// there's probably a better way to do this using interfaces
func (a ArticleList) renderTemplate(w http.ResponseWriter, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "base", a)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}

func (a Article) renderTemplate(w http.ResponseWriter, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "base", a)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
func (a *Article) loadContent() error {
	contentDir := "./blog/articles/"
	fileName := contentDir + a.Content.Path
	var err error
	a.Content.Body, err = os.ReadFile(fileName)
	if err != nil {
		return err
	}
	a.Content.BodyHTML = template.HTML(a.Content.Body)
	return nil
}
