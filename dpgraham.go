package main

import (
	"errors"
	"html/template"
	"strings"

	// "io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title    string
	Body     []byte
	Path     string
	Articles []string
}

func loadArticle(title string) (*Page, error) {
	htmlDir := "./html/"
	filename := htmlDir + title + ".txt"
	body, _ := os.ReadFile(filename)
	pageContent := Page{
		Title: title,
		Body:  body,
		Path:  htmlDir,
	}
	return &pageContent, nil
}

func (p *Page) readArticles() error {
	dir := "./articles/"
	f, _ := os.Open(dir)
	files, _ := f.ReadDir(0)
	for _, v := range files {
		articleTitle := convertTitles(v.Name())
		p.Articles = append(p.Articles, articleTitle)
	}
	return nil
}

func convertTitles(filename string) string {
	filename = strings.ReplaceAll(filename, ".txt", "")
	titleParts := strings.Split(filename, "_")
	// fmt.Printf("%T\n", titleParts)
	title := strings.Join(titleParts, " ")
	return title
}

var templatePaths = []string{
	"./html/index.html",
	"./html/list.html",
}

var templates = template.Must(template.ParseFiles(templatePaths...))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadArticle("home")
	renderTemplate(w, "index", p)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/blog" {
		p, _ := loadArticle("home")
		p.readArticles()
		renderTemplate(w, "list", p)
	} else {
		p, _ := loadArticle("home")
		renderTemplate(w, "list", p)
	}
}

func main() {
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog", blogHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
