package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teixy/go/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT},
	}))

	e.GET("/books/all", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBook)

	e.POST("/books", controllers.CreateBook)

	e.PUT("/books/:id", controllers.UpdateBook)

	e.Logger.Fatal(e.Start(":8080"))
}
