package route

import (
    "github.com/labstack/echo/v4"
    "openai/controller"
	"openai/usecase"
)

func SetLaptopRecomendedRoutes(e *echo.Echo) {
    laptopRecomendedUsecase := usecase.NewUseCase()
    laptopRecomendedController := controller.NewLaptopRecomendedController(laptopRecomendedUsecase)

    e.POST("/laptop", laptopRecomendedController.GetLaptopRecomended)
}
