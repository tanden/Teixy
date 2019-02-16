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

func GetAllArticles(limit int, offset int) Articles {

	data := db.CreateConectionTeixyArticle()

	result := Articles{}
	query := "SELECT * FROM articles WHERE id BETWEEN ? and ?"
	err := data.Select(&result.Articles, query, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
