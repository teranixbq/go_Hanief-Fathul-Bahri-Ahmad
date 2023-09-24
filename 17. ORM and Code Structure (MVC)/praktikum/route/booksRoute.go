package route

import (
	"github.com/labstack/echo"
	"praktikum/rest/controller"
)

func BooksRoute(e *echo.Echo) {
	books := e.Group("/books")
	books.GET("", controller.GetAllBooksController)
	books.GET("/:id", controller.GetBooksByIdController)
	books.POST("", controller.CreateBooksController)
	books.DELETE("/:id", controller.DeleteBooksByIdController)
	books.PUT("/:id", controller.UpdateBooksController)
}
