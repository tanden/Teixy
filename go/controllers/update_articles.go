package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type Status struct {
	Status int `validate:"required,min=1,numeric"`
}

type Params struct {
	Id
	Article
	Status
}

func UpdateArticle(c echo.Context) error {
	article_id, _ := strconv.Atoi(c.Param("id"))
	max_score, _ := strconv.Atoi(c.FormValue("max_score"))
	min_score, _ := strconv.Atoi(c.FormValue("min_score"))
	status, _ := strconv.Atoi(c.FormValue("status")) 
	params := Params{
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
