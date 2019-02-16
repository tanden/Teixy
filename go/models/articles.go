package models

import (
	_ "github.com/lib/pq"
	"github.com/teixy/go/db"
	"log"
)

type Article struct {
	Id         int    `json:"id"`
	Max_Score  int    `json:"max_score"`
	Min_Score  int    `json:"min_score"`
	Title      string `json:"title"`
	Punch_Line string `json:"punch_line"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	Mtime      string `json:"mtime"`
	Ctime      string `json:"ctime"`
}

type Articles struct {
	Articles []Article `json:"article"`
}

func GetAllArticles(min_id int, max_id int) Articles {

	data := db.CreateConectionTeixyArticles()

	result := Articles{}
	query := "SELECT * FROM articles WHERE id BETWEEN ? and ?"
	err := data.Select(&result.Articles, query, min_id, max_id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetArticle(id int) []Article {

	data := db.CreateConectionTeixyArticles()

	var result []Article
	query := "SELECT * FROM articles WHERE id = ?"
	err := data.Select(&result, query, id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
