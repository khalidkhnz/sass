package main

import (
	"net/http"

	"github.com/khalidkhnz/sass/go-ecom/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitEnv()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		 return c.JSON(http.StatusOK,map[string]any{
			"success":true,
			"message":"Hello From Go-Ecom",
		 })
	})
	e.Logger.Fatal(e.Start(config.GetPort()))
}