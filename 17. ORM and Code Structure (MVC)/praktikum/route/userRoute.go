package route

import (
	"github.com/labstack/echo"
	"praktikum/rest/controller"
)

func UserRoute(e *echo.Echo) {
	
	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.DELETE("/:id", controller.DeleteUserController)
	user.PUT("/:id", controller.UpdateUserController)
	
}
