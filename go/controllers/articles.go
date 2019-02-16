package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type Params struct {
	Limit  int `validate:"required,min=1,numeric"`
	Offset int `validate:"required,min=1,numeric"`
}

var validate *validator.Validate

func paramsValidator(c echo.Context) (error, *Params) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))
	validate = validator.New()
	params := &Params{
		Limit:  limit,
		Offset: offset,
	}
	err := validate.Struct(params)
	return err, params
}

func GetAllArticles(c echo.Context) error {

	err, params := paramsValidator(c)
	if err != nil {
		fmt.Println("err", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Parameter")
	}

	result := models.GetAllArticles(params.Limit, params.Offset)
	return c.JSON(http.StatusOK, result)
}
