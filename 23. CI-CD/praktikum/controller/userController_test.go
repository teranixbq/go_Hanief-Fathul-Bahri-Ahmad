package controller

import (
	"clean/dto"
	"clean/middleware"
	"clean/mocks"
	"clean/model"
	"encoding/json"
	"errors"

	// "errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UsersResponseSuccess struct {
	Message string
	Data    []model.User
}

type ResponseGlobal struct {
	Status string
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	mockUserUsecase := new(mocks.UserUsecase)
	controller := NewUserController(mockUserUsecase)

	t.Run("Succes Create User", (func(t *testing.T) {
		reqbody := `{
		"email": "hanief@gmail.com", 
		"password":"1234gs"}`

		mockUserUsecase.On("Create", mock.AnythingOfType("dto.CreateUserRequest")).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, controller.Create(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}

	}))
}

func TestGetAll(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UserUsecase)
	returnData := []model.User{{ID: 1, Email: "alta@mail.id", Password: "sdad"}}

	t.Run("Success Get All", func(t *testing.T) {
		usecase.On("SelectAll").Return(returnData, nil).Once()
		token, errToken := middleware.CreateToken(returnData[0].ID)
		if errToken != nil {
			assert.NoError(t, errToken)
		}
		srv := NewUserController(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		if assert.NoError(t, middleware.JWTMiddleware()(echo.HandlerFunc(srv.SelectAll))(echoContext)) {
			responseBody := rec.Body.String()
			responseData := []model.User{}
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].ID, responseData[0].ID)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		expectedErr := errors.New("error message")
		usecase.On("SelectAll").Return([]model.User{}, expectedErr).Once()
		token, errToken := middleware.CreateToken(1)
		if errToken != nil {
			assert.NoError(t, errToken)
		}
		srv := NewUserController(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		if assert.NoError(t, middleware.JWTMiddleware()(echo.HandlerFunc(srv.SelectAll))(echoContext)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		}
		usecase.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	e := echo.New()
	mockUserUsecase := new(mocks.UserUsecase)
	t.Run("Succes Login", func(t *testing.T) {
		
		controller := NewUserController(mockUserUsecase)
		mockUserUsecase.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(model.User{}, "token", nil)

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, controller.Login(c)) {
			bodyResponses := rec.Body.String()
			var user ResponseGlobal
			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success", user.Status)

		}
	})

	t.Run("failed login ", func(t *testing.T) {
		mockUserUsecase := new(mocks.UserUsecase)
		expectedErr := errors.New("Login failed")
		mockUserUsecase.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(model.User{}, "", expectedErr)

		controller := NewUserController(mockUserUsecase)

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, controller.Login(c)) {
			bodyResponses := rec.Body.String()
			var user dto.ResponseLogin

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	mockUserUsecase.AssertExpectations(t)

}
