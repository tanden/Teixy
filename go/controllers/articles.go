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

type Id struct {
	Id int `validate:"required,min=1,numeric"`
}

type Status struct {
	Status int `validate:"min=0,max=1,numeric"`
}

type Article struct {
	Min_Score  int    `validate:"required,min=1,numeric"`
	Max_Score  int    `validate:"required,min=1,numeric,gtefield=Min_Score"`
	Title      string `validate:"required"`
	Punch_Line string `validate:"required"`
	Content    string `validate:"required"`
}

type UpdateParams struct {
	Id
	Article
	Status
}

type QueryResult struct {
	LastInsertId int64
	RowAffected  int64
}

func GetAllArticles(c echo.Context) error {
	min_id, _ := strconv.Atoi(c.QueryParam("min_id"))
	max_id, _ := strconv.Atoi(c.QueryParam("max_id"))
	params := MinMax{min_id, max_id}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := models.GetAllArticles(params.Min_Id, params.Max_Id)
	return c.JSON(http.StatusOK, result)
}

func GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	params := Id{id}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := models.GetArticle(id)
	return c.JSON(http.StatusOK, result)
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

func UpdateArticle(c echo.Context) error {
	article_id, _ := strconv.Atoi(c.Param("id"))
	max_score, _ := strconv.Atoi(c.FormValue("max_score"))
	min_score, _ := strconv.Atoi(c.FormValue("min_score"))
	status, _ := strconv.Atoi(c.FormValue("status")) 
	params := UpdateParams{
		Id{article_id},
		Article{
			min_score,
			max_score,
			c.FormValue("title"),
			c.FormValue("punch_line"),
			c.FormValue("content"),
		},
		Status{status},
	}

	validate = validator.New()
	err := validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := models.UpdateArticle(
		params.Id.Id,
		params.Min_Score,
		params.Max_Score,
		params.Title,
		params.Punch_Line,
		params.Content,
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
