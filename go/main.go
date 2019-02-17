package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teixy/go/controllers"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World!")
	})

	e.GET("/books/all", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBook)

	e.POST("/books", controllers.CreateBook)

	e.PUT("/books/:id", controllers.UpdateBook)

	e.Logger.Fatal(e.Start(":8080"))
}
