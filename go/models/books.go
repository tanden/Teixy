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
	Book_Id    int    `json:"book_id"`
	Isbn       uint64 `json:"isbn"`
	Min_Score  int    `json:"min_score"`
	Max_Score  int    `json:"max_score"`
	Title      string `json:"title"`
	Punch_Line string `json:"punch_line"`
	Article    string `json:"article"`
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
	query := "SELECT * FROM books WHERE book_id BETWEEN ? and ?"
	err := Data.Select(&result.Books, query, min_id, max_id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func GetBook(book_id int) []Book {

	var result []Book
	query := "SELECT * FROM books WHERE book_id = ?"
	err := Data.Select(&result, query, book_id)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func CreateBook(isbn uint64, min_score int, max_score int, title string, punch_line string, article string) (sql.Result, error) {

	stmt, err := Data.Prepare(`
	INSERT INTO books (
		isbn,
		min_score,
		max_score,
		title,
		punch_line,
		article,
		status
	) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return stmt.Exec(
		isbn,
		min_score,
		max_score,
		title,
		punch_line,
		article,
		StatusOff,
	)
}

func UpdateBook(book_id int, min_score int, max_score int, title string, punch_line string, article string, status int) (sql.Result, error) {

	stmt, err := Data.Prepare(`
	UPDATE books SET
		min_score = ?,
		max_score = ?,
		title = ?,
		punch_line = ?,
		article = ?,
		status = ?
	WHERE book_id = ?
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
		article,
		status,
		book_id,
	)
}
