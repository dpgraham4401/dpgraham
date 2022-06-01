package views

import "html/template"

var Index = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/index.gohtml"))

var Blogs = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/blog.gohtml"))

var Blog = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/article.gohtml"))

var Error = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/error.gohtml"))

var About = template.Must(template.ParseFiles(
	"./views/base.gohtml",
	"./views/about.gohtml"))
