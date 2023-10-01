package route

import (
	"github.com/labstack/echo/v4"
	"praktikum/rest/controller"

	
)

func UserRoute(e *echo.Echo) {
	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.DELETE("/:id", controller.DeleteUserController)
	user.PUT("/:id", controller.UpdateUserController)
	
	
}
