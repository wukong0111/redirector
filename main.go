package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
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
	e.Logger.Fatal(e.Start(":3000"))
}
