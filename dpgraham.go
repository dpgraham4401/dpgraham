package main

import (
	"fmt"
	"html/template"
	"strings"

	// "io/fs"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title    string
	Body     []byte
	Path     string
	LinkList []Link
}

type Link struct {
	Article string
	Link    string
}

func loadArticle(title string) (*Page, error) {
	htmlDir := "./articles/"
	filename := htmlDir + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error opening ", filename)
	}
	pageContent := Page{
		Title: title,
		Body:  body,
		Path:  htmlDir,
	}
	return &pageContent, nil
}

func (p *Page) readArticleList() error {
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
	filename = strings.ReplaceAll(filename, ".txt", "")
	titleParts := strings.Split(filename, "_")
	title := strings.Join(titleParts, " ")
	return title, filename
}

var templatePaths = []string{
	"./html/index.html",
	"./html/blog_home.html",
	"./html/article.html",
}

var templates = template.Must(template.ParseFiles(templatePaths...))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("invalid Page Title")
// 	}
// 	return m[2], nil
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadArticle("home")
	renderTemplate(w, "index", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog/" {
		pBlank := Page{}
		p := &pBlank
		p.readArticleList()
		renderTemplate(w, "blog_home", p)
	} else {
		title := r.URL.Path[len("/blog/"):]
		p, _ := loadArticle(title)
		renderTemplate(w, "article", p)
	}
}

func main() {
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog/", blogHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
