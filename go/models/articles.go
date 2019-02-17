package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/teixy/go/db"
	"log"
)

const (
	StatusOff = 0
	StatusOn  = 1
)

type Book struct {
	Id         int    `json:"id"`
	Min_Score  int    `json:"min_score"`
	Max_Score  int    `json:"max_score"`
	Title      string `json:"title"`
	Punch_Line string `json:"punch_line"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	Mtime      string `json:"mtime"`
	Ctime      string `json:"ctime"`
}

type Books struct {
	Books []Book `json:"books"`
}

var Data *sqlx.DB

func init() {
	Data = db.CreateConectionTeixyBooks()
}

func GetAllBooks(min_id int, max_id int) Books {

	result := Books{}
	query := "SELECT * FROM books WHERE id BETWEEN ? and ?"
	err := Data.Select(&result.Books, query, min_id, max_id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetBook(book_id int) []Book {

	var result []Book
	query := "SELECT * FROM books WHERE id = ?"
	err := Data.Select(&result, query, book_id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func CreateBook(min_score int, max_score int, title string, punch_line string, content string) (sql.Result, error) {

	stmt, err := Data.Prepare(`
	INSERT INTO books (
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

func UpdateBook(book_id int, min_score int, max_score int, title string, punch_line string, content string, status int) (sql.Result, error) {

	stmt, err := Data.Prepare(`
	UPDATE books SET
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
		book_id,
	)
}
