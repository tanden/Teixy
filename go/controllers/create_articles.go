package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type Article struct {
	Min_Score  int    `validate:"required,min=1,numeric"`
	Max_Score  int    `validate:"required,min=1,numeric,gtefield=Min_Score"`
	Title      string `validate:"required"`
	Punch_Line string `validate:"required"`
	Content    string `validate:"required"`
}

type QueryResult struct {
	LastInsertId int64
	RowAffected  int64
}

func CreateArticle(c echo.Context) error {

	min_score, _ := strconv.Atoi(c.FormValue("min_score"))
	max_score, _ := strconv.Atoi(c.FormValue("max_score"))
	params := Article{
		min_score,
		max_score,
		c.FormValue("title"),
		c.FormValue("punch_line"),
		c.FormValue("content"),
	}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := models.CreateArticle(
		params.Min_Score,
		params.Max_Score,
		params.Title,
		params.Punch_Line,
		params.Content,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()
	query_result := QueryResult{id, rows}

	return c.JSON(http.StatusCreated, query_result)
}
