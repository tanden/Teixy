package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var validate *validator.Validate

// form validation
type MinMax struct {
	Min_Id int `validate:"required,min=1,numeric"`
	Max_Id int `validate:"required,min=1,numeric,gtefield=Min_Id"`
}

type BookId struct {
	Id int `validate:"required,min=1,numeric"`
}

type Score struct {
	Min_Score int `validate:"required,min=1,max=9999,numeric"`
	Max_Score int `validate:"required,min=1,max=10000,numeric,gtefield=Min_Score"`
}

type Book struct {
	Isbn       int    `validate:"required,min=9780000000000,max=9800000000000,numeric"`
	Title      string `validate:"required"`
	Punch_Line string `validate:"required"`
	Article    string `validate:"required"`
}

type Status struct {
	Status int `validate:"min=0,max=1,numeric"`
}

type UpdateParams struct {
	BookId
	Book
	Status
}

type UpdateParamsWithoutScore struct {
	BookId
	Book
	Status
}

type QueryResult struct {
	LastInsertId int64
	RowAffected  int64
}

func GetAllBooks(c echo.Context) error {
	min_id, _ := strconv.Atoi(c.QueryParam("min_id"))
	max_id, _ := strconv.Atoi(c.QueryParam("max_id"))
	params := MinMax{min_id, max_id}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := models.GetAllBooks(params.Min_Id, params.Max_Id)
	return c.JSON(http.StatusOK, result)
}

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	params := BookId{id}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := models.GetBook(params.Id)
	return c.JSON(http.StatusOK, result)
}

func CreateBook(c echo.Context) error {

	isbn, _ := strconv.Atoi(c.FormValue("isbn"))

	min_score := sql.NullInt64{0, false}
	max_score := sql.NullInt64{0, false}

	if c.FormValue("min_score") != "" && c.FormValue("max_score") != "" {
		min_score_value, _ := strconv.Atoi(c.FormValue("min_score"))
		max_score_value, _ := strconv.Atoi(c.FormValue("max_score"))
		scores := Score{
			min_score_value,
			max_score_value,
		}
		validate = validator.New()
		err := validate.Struct(scores)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		min_score = sql.NullInt64{int64(min_score_value), true}
		max_score = sql.NullInt64{int64(max_score_value), true}
	}

	book := Book{
		isbn,
		c.FormValue("title"),
		c.FormValue("punch_line"),
		c.FormValue("article"),
	}

	validate = validator.New()
	err := validate.Struct(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := models.CreateBook(
		book.Isbn,
		min_score,
		max_score,
		book.Title,
		book.Punch_Line,
		book.Article,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()
	query_result := QueryResult{id, rows}

	return c.JSON(http.StatusCreated, query_result)
}

func UpdateBook(c echo.Context) error {

	min_score := sql.NullInt64{0, false}
	max_score := sql.NullInt64{0, false}

	if c.FormValue("min_score") != "" && c.FormValue("max_score") != "" {
		min_score_value, _ := strconv.Atoi(c.FormValue("min_score"))
		max_score_value, _ := strconv.Atoi(c.FormValue("max_score"))
		scores := Score{
			min_score_value,
			max_score_value,
		}
		validate = validator.New()
		err := validate.Struct(scores)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		min_score = sql.NullInt64{int64(min_score_value), true}
		max_score = sql.NullInt64{int64(max_score_value), true}
	}

	book_id, _ := strconv.Atoi(c.Param("id"))
	status, _ := strconv.Atoi(c.FormValue("status"))
	isbn, _ := strconv.Atoi(c.FormValue("isbn"))

	params := UpdateParams{
		BookId{book_id},
		Book{
			isbn,
			c.FormValue("title"),
			c.FormValue("punch_line"),
			c.FormValue("article"),
		},
		Status{status},
	}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := models.UpdateBook(
		params.Id,
		params.Isbn,
		min_score,
		max_score,
		params.Title,
		params.Punch_Line,
		params.Article,
		params.Status.Status,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	update_id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()
	query_result := QueryResult{update_id, rows}

	return c.JSON(http.StatusCreated, query_result)
}
