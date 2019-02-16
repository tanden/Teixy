package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var validate *validator.Validate

type MinMaxParams struct {
	Min_Id int `validate:"required,min=1,numeric"`
	Max_Id int `validate:"required,min=1,numeric,gtefield=Min_Id"`
}

func (params *MinMaxParams) Validate() error {
	validate = validator.New()
	err := validate.Struct(params)
	return err
}

func GetAllArticles(c echo.Context) error {

	min_id, _ := strconv.Atoi(c.QueryParam("min_id"))
	max_id, _ := strconv.Atoi(c.QueryParam("max_id"))

	params := MinMaxParams{min_id, max_id}
	err := params.Validate()
	if err != nil {
		fmt.Println("err", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Parameter")
	}

	result := models.GetAllArticles(params.Min_Id, params.Max_Id)
	return c.JSON(http.StatusOK, result)
}
