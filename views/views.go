package views

import "html/template"

var Index = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/index.gohtml"))

var Blogs = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/blog_home.gohtml"))

var Blog = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/article.gohtml"))
