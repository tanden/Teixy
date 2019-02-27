package controllers

import (
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

type Status struct {
	Status int `validate:"min=0,max=1,numeric"`
}

type Book struct {
	Isbn       uint64 `validate:"required,min=9780000000000,numeric"`
	Min_Score  int    `validate:"required,min=1,numeric"`
	Max_Score  int    `validate:"required,min=1,numeric,gtefield=Min_Score"`
	Title      string `validate:"required"`
	Punch_Line string `validate:"required"`
	Article    string `validate:"required"`
}

type UpdateParams struct {
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

	min_score, _ := strconv.Atoi(c.FormValue("min_score"))
	max_score, _ := strconv.Atoi(c.FormValue("max_score"))
	isbn, _ := strconv.ParseUint(c.FormValue("isbn"), 10, 64)
	params := Book{
		isbn,
		min_score,
		max_score,
		c.FormValue("title"),
		c.FormValue("punch_line"),
		c.FormValue("article"),
	}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := models.CreateBook(
		params.Isbn,
		params.Min_Score,
		params.Max_Score,
		params.Title,
		params.Punch_Line,
		params.Article,
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
	book_id, _ := strconv.Atoi(c.Param("id"))
	max_score, _ := strconv.Atoi(c.FormValue("max_score"))
	min_score, _ := strconv.Atoi(c.FormValue("min_score"))
	status, _ := strconv.Atoi(c.FormValue("status")) 
	params := UpdateParams{
		BookId{book_id},
		Book{
			min_score,
			max_score,
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
		params.Min_Score,
		params.Max_Score,
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
