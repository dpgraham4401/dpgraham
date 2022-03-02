package main

import (
	"html/template"
	"strings"

	"log"
	"net/http"
	"os"
)

var templatePaths = []string{
	"./html/index.html",
	"./html/blog_home.html",
	"./html/article.html",
	"./articles/home.html",
	"./articles/first_post.html",
	"./articles/too_much_time_on_nvim.html",
	"./articles/another_articles.html",
}

var templates = template.Must(template.ParseFiles(templatePaths...))

// Article for all things webpage
type Article struct {
	Title string
	Body  []byte
	Path  string
}

// ArticleList represents any page used for basic navigation
type ArticleList struct {
	LinkList []Link
}

// Link to be used in article[]
type Link struct {
	Article string
	Link    string
}

// Page interface includes methods for all types of page content
type Page interface {
	loadContent()
	renderTemplate()
}

func (article *Article) loadContent(title string) error {
	htmlDir := "./articles/"
	filename := htmlDir + title + ".html"
	body, _ := os.ReadFile(filename)
	article.Body = body
	article.Title = title
	article.Path = htmlDir
	return nil
}

func (articleList *ArticleList) loadContent() error {
	dir := "./articles/"
	f, _ := os.Open(dir)
	files, _ := f.ReadDir(0)
	for _, v := range files {
		articleTitle, link := convertTitles(v.Name())
		articleList.LinkList = append(articleList.LinkList, Link{Article: articleTitle, Link: link})
	}
	return nil
}

func convertTitles(filename string) (string, string) {
	filename = strings.ReplaceAll(filename, ".html", "")
	titleParts := strings.Split(filename, "_")
	title := strings.Join(titleParts, " ")
	return title, filename
}

func (article *Article) renderTemplate(w http.ResponseWriter, tmpl string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, tmpl+".html", article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (articleList *ArticleList) renderTemplate(w http.ResponseWriter, tmpl string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, tmpl+".html", articleList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog/" {
		var content ArticleList
		content.loadContent()
		content.renderTemplate(w, "blog_home")
	} else {
		var content Article
		content.loadContent(r.URL.Path[len("/blog/"):])
		content.renderTemplate(w, "article")
	}
}

func main() {
	// Static files and assets
	fs := http.FileServer(http.Dir("html/assets/"))
	mux := http.NewServeMux()
	mux.Handle("/html/assets/", http.StripPrefix("/html/assets/", fs))

	// Routing
	mux.HandleFunc("/blog/", blogHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
