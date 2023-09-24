// +build tools

package main

import (
	"praktikum/rest/config"
	"praktikum/rest/route"
	"github.com/labstack/echo"
)

// ---------------------------------------------------
func main() {
	config.InitDB()
	e := echo.New()
	route.UserRoute(e)
	route.BooksRoute(e)
	route.BlogsRoute(e)
	e.Logger.Fatal(e.Start(":8000"))
}
