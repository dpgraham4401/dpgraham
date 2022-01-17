package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "html/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("home")
	t, _ := template.ParseFiles("./html/home.html")
	t.Execute(w, p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
