package route

import (
	"github.com/labstack/echo/v4"
	"praktikum/rest/controller"
	"praktikum/rest/middleware"
)

func BooksRoute(e *echo.Echo) {

	books := e.Group("/books",middleware.JWTMiddleware())
	books.GET("", controller.GetAllBooksController)
	books.GET("/:id", controller.GetBooksByIdController)
	books.POST("", controller.CreateBooksController)
	books.DELETE("/:id", controller.DeleteBooksByIdController)
	books.PUT("/:id", controller.UpdateBooksController)
}
