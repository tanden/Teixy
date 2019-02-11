package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/teixy/go/models"
	"net/http"
)

func GetArticles(c echo.Context) error {
	result := models.GetArticles()
	return c.JSON(http.StatusOK, result)
}
