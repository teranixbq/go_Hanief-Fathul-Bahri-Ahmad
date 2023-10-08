package controller

import (
	"clean/dto"
	"clean/helper"
	"clean/middleware"
	"clean/usercase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	useCase usercase.UserUsecase
}

func NewUserController(userUsecase usercase.UserUsecase) *userController {
	return &userController{useCase: userUsecase}
}

func (uc *userController) Create(e echo.Context) error {
	var payload dto.CreateUserRequest

	if err := e.Bind(&payload); err != nil {
		return err
	}

	err := uc.useCase.Create(payload)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Succes Create User"))
}

func (uc *userController) SelectAll(c echo.Context) error {
	id := middleware.ExtractToken(c)

	if id == 0 {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}
	users, err := uc.useCase.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    dto.FromAllData(users),
	})
}

func (uc *userController) Login(e echo.Context) error {
	var payload dto.CreateUserRequest

	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Request"))
	}

	user, token, err := uc.useCase.Login(payload.Email, payload.Password)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	response := dto.ResponseLogin{

		Email: user.Email,
		Token: token,
	}

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Succes Login", response))
}
