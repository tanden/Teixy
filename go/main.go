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

	e.GET("/articles/all", controllers.GetAllArticles)
	e.GET("/articles/:id", controllers.GetArticle)

	e.POST("/articles", controllers.CreateArticle)

	e.PUT("/articles/:id", controllers.UpdateArticle)

	e.Logger.Fatal(e.Start(":8080"))
}
