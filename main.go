package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

//Db holds database connection
var Db *sql.DB

func init() {
	var err error
	pgConn := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"))
	Db, err = sql.Open("postgres", pgConn)
	CheckError(err)
}

// Article captures metadata about a blog post or tutorial etc.
type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	LastUpdate  string `json:"updateDate"`
	CreateDate  string `json:"createDate"`
	Published   bool   `json:"publish"`
	ArticleType string `json:"type"`
	Content     string `json:"content"`
}

func main() {
	_, blog := getBlog(1)
	fmt.Println(blog.Title)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBlog(id int) (error, Article) {
	blog := Article{}
	err := Db.QueryRow("SELECT * FROM article WHERE id = $1", id).Scan(
		&blog.Id, &blog.Title, &blog.LastUpdate, &blog.Published,
		&blog.ArticleType, &blog.CreateDate, &blog.Content)
	if err != nil {
		fmt.Println(err)
		return nil, blog
	}
	return nil, blog
}
