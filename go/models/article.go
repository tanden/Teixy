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

func GetArticles() Articles {

	data := db.CreateConectionTeixyArticle()

	result := Articles{}
	query := "SELECT * FROM articles ORDER BY id"
	err := data.Select(&result.Articles, query)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
