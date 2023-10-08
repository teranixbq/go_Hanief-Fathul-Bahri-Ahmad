package route

import (
	"clean/controller"
	"clean/middleware"
	"clean/repository"
	"clean/usercase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute (e *echo.Echo,db *gorm.DB){
	userRepository := repository.NewUserRepository(db)
	userService := usercase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userService)

	e.POST("/users", userController.Create)
	e.GET("/users", userController.SelectAll,middleware.JWTMiddleware())
	e.POST("/login",userController.Login)
}