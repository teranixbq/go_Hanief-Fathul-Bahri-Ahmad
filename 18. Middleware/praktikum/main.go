// +build tools

package main

import (
	"praktikum/rest/config"
	"praktikum/rest/route"
	"github.com/labstack/echo/v4"
	m "praktikum/rest/middleware"
)

// ---------------------------------------------------
func main(){
	e := echo.New() 
	config.InitDB()
	route.UserRoute(e)
	route.BooksRoute(e)
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
