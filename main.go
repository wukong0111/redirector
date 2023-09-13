package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		code := c.QueryParam("code")
		state := c.QueryParam("state")
		path := c.QueryParam("path")
		u, err := url.Parse(path)
		if err != nil {
			return err
		}
		params := url.Values{}
		params.Add("code", code)
		params.Add("state", state)
		u.RawQuery = params.Encode()
		log.Println(path, u.String())
		return c.Redirect(http.StatusTemporaryRedirect, u.String())
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)
	e.Logger.Fatal(e.Start(address))
}
