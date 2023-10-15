package controller

import (
	"log"
	"net/http"
	"openai/dto"
	"openai/usecase"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type LaptopRecomendedController struct {
	usecase usecase.LaptopRecomendedUsecase
}

func NewLaptopRecomendedController(usecase usecase.LaptopRecomendedUsecase) *LaptopRecomendedController {
	return &LaptopRecomendedController{usecase: usecase}
}

func (controller *LaptopRecomendedController) GetLaptopRecomended(c echo.Context) error {
	var request dto.RequestRecomended
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	answer, err := controller.usecase.LaptopRecomended(request, os.Getenv("OPEN_API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Adding "answer" to the response
	responseDTO := dto.ResponseRecomended{
		Status:         "Success",
		Recommendation: answer,
	}

	return c.JSON(http.StatusOK, responseDTO)
}
