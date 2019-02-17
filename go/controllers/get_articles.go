package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var validate *validator.Validate

type MinMax struct {
	Min_Id int `validate:"required,min=1,numeric"`
	Max_Id int `validate:"required,min=1,numeric,gtefield=Min_Id"`
}

type Id struct {
	Id int `validate:"required,min=1,numeric"`
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
