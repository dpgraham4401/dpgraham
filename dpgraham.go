package main

import (
	"html/template"
	"strings"

	"log"
	"net/http"
	"os"
)

// Page interface includes methods for all types of page content
type Page interface {
	LoadArticle()
	renderTemplate()
}

// Article for all things webpage
type Article struct {
	Title string
	Body  []byte
	Path  string
}

// AricleList represents any page used for basic navigation
type ArticleList struct {
	LinkList []Link
}

// Link to be used in article[]
type Link struct {
	Article string
	Link    string
}

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

func (Article *Article) loadArticle(title string) error {
	htmlDir := "./articles/"
	filename := htmlDir + title + ".html"
	body, _ := os.ReadFile(filename)
	Article.Body = body
	// pageContent := Article{
	// 	Title: title,
	// 	Body:  body,
	// 	Path:  htmlDir,
	// }
	// return &pageContent, nil
}

func (p *ArticleList) readArticleList() error {
	dir := "./articles/"
	f, _ := os.Open(dir)
	files, _ := f.ReadDir(0)
	for _, v := range files {
		articleTitle, link := convertTitles(v.Name())
		p.LinkList = append(p.LinkList, Link{Article: articleTitle, Link: link})
	}
	return nil
}

func convertTitles(filename string) (string, string) {
	filename = strings.ReplaceAll(filename, ".html", "")
	titleParts := strings.Split(filename, "_")
	title := strings.Join(titleParts, " ")
	return title, filename
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Article) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog/" {
		var articleList ArticleList
		// p := &articleList
		articleList.readArticleList()
		// renderTemplate(w, "blog_home", &articleList)
	} else {
		title := r.URL.Path[len("/blog/"):]
		p, _ := loadArticle(title)
		renderTemplate(w, "article", p)
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
