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
	Min_Id int `validate:"required,min=1,numeric"`
	Max_Id int `validate:"required,min=1,numeric,gtefield=Min_Id"`
}

var validate *validator.Validate

func paramsValidator(c echo.Context) (error, *Params) {
	min_id, _ := strconv.Atoi(c.QueryParam("min_id"))
	max_id, _ := strconv.Atoi(c.QueryParam("max_id"))
	validate = validator.New()
	params := &Params{
		Min_Id: min_id,
		Max_Id: max_id,
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

	result := models.GetAllArticles(params.Min_Id, params.Max_Id)
	return c.JSON(http.StatusOK, result)
}
