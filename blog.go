package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

var templatePaths = []string{
	"./templates/index.html",
	"./templates/blog_home.html",
	"./templates/article.html",
}

var templates = template.Must(template.ParseFiles(templatePaths...))

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
func loadArticles() ArticleList {
	var allArticles ArticleList
	articlePath := "./blog/entries/"
	articleDir, _ := os.Open(articlePath)
	files, _ := articleDir.ReadDir(0)
	for _, v := range files {
		var newArticle Article
		file, _ := ioutil.ReadFile(articlePath + v.Name())
		_ = json.Unmarshal([]byte(file), &newArticle)
		allArticles.Articles = append(allArticles.Articles, newArticle)
	}
	sort.Slice(allArticles.Articles, func(i, j int) bool {
		return allArticles.Articles[i].Id > allArticles.Articles[j].Id
	})
	return allArticles
}

func (a ArticleList) renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a Content) renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *Article) loadContent() {
	contentDir := "./blog/articles/"
	fileName := contentDir + a.Content.Path
	a.Content.Body, _ = os.ReadFile(fileName)
	a.Content.BodyHTML = template.HTML(a.Content.Body)
}
