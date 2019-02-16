package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"net/http"
	"strconv"
)

func GetAllArticles(c echo.Context) error {

	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))

	result := models.GetAllArticles(limit, offset)
	return c.JSON(http.StatusOK, result)
}
