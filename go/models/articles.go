package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/teixy/go/db"
	"log"
)

const StatusOff = 0
const StatusOn = 1

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

func CreateArticle(min_score int, max_score int, title string, punch_line string, content string) (sql.Result, error) {

	data := db.CreateConectionTeixyArticles()

	stmt, err := data.Prepare(`
	INSERT INTO articles (
		min_score,
		max_score,
		title,
		punch_line,
		content,
		status
	) 
	VALUES (?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return stmt.Exec(
		min_score,
		max_score,
		title,
		punch_line,
		content,
		StatusOff,
	)
}

func UpdateArticle(id int, min_score int, max_score int, title string, punch_line string, content string, status int) (sql.Result, error) {

	data := db.CreateConectionTeixyArticles()

	stmt, err := data.Prepare(`
	UPDATE articles SET
		min_score = ?,
		max_score = ?,
		title = ?,
		punch_line = ?,
		content = ?,
		status = ?
	WHERE id = ?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return stmt.Exec(
		min_score,
		max_score,
		title,
		punch_line,
		content,
		status,
		id,
	)
}
