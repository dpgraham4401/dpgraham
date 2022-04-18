package views

import "html/template"

var Index = template.Must(template.ParseFiles(
	"./views/base.gotmpl",
	"./views/index.gotmpl"))

var Blogs = template.Must(template.ParseFiles(
	"./views/base.gotmpl",
	"./views/blog_home.gotmpl"))

var Blog = template.Must(template.ParseFiles(
	"./views/base.gotmpl",
	"./views/article.gotmpl"))
