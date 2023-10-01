package controller

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"net/http"
	"net/http/httptest"
	"praktikum/rest/config"
	"praktikum/rest/controller/responsejson"
	"praktikum/rest/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "strings"
)

type testCases struct {
	path       string
	expectCode int
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InitEchoTestAPINoDop() *echo.Echo {
	config.InitDBTestNoDrop()
	e := echo.New()
	return e
}

var (
	data_user = model.User{
		Name:     "alta",
		Email:    "alta@gmail.com",
		Password: "12345",
	}
)

func InsertMockDataUser() error {
	var err error
	if err = config.DB.Create(&data_user).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateUserController(t *testing.T) {
	test := testCases{
		path: "/users",
	}
	t.Run("Success Create", func(t *testing.T) {

		test.expectCode = http.StatusOK

		e := InitEchoTestAPI()
		body, err := json.Marshal(data_user)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPost, test.path, bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {

			bodyResponses := rec.Body.String()

			var user responsejson.ResponCreate
			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.expectCode, rec.Code)
			assert.Equal(t, data_user.Name, user.Data.Name)
			assert.Equal(t, data_user.Email, user.Data.Email)
			assert.Equal(t, "success", user.Status)
		}
	})

	t.Run("failed Create", func(t *testing.T) {

		test.expectCode = http.StatusBadRequest

		e := InitEchoTestAPI()
		config.DB.Migrator().DropTable(&model.User{})

		req := httptest.NewRequest(http.MethodPost, test.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call function on controller
		if assert.NoError(t, CreateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user responsejson.ResponCreate

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, test.expectCode, rec.Code)
			assert.Equal(t, "failed", user.Status)
		}
	})

}

func TestGetUsersController(t *testing.T) {
	test := testCases{
		path: "/users",
	}
	t.Run("Success Get All User", func(t *testing.T) {

		test.expectCode = http.StatusOK

		e := InitEchoTestAPI()
		InsertMockDataUser()
		req := httptest.NewRequest(http.MethodGet, test.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetUsersController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}

	})

	t.Run("Failed Get All User", func(t *testing.T) {
		test.expectCode = http.StatusInternalServerError

		e := InitEchoTestAPI()
		config.DB.Migrator().DropTable(&model.User{})

		req := httptest.NewRequest(http.MethodPost, test.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call function on controller
		if assert.NoError(t, GetUsersController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}

	})
}

func TestGetUserByIdController(t *testing.T) {
	e := InitEchoTestAPI()
	test := testCases{
		path: "/users/:id",
	}

	t.Run("Succes GetUserById", func(t *testing.T) {
		test.expectCode = http.StatusOK

		e := InitEchoTestAPI()
		InsertMockDataUser()
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(test.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetUserController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}
	})

	t.Run("Failed GetUserById", func(t *testing.T) {
		test.expectCode = http.StatusBadRequest

		InsertMockDataUser()
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("#")

		if assert.NoError(t, GetUserController(c)) {

			assert.Equal(t, test.expectCode, res.Code)

		}
	})

}

func TestUpdateUserController(t *testing.T) {
	e := InitEchoTestAPINoDop()

	var (
		data_user = model.User{
			Name:     "hanief",
			Email:    "hanief@gmail.com",
			Password: "12345",
		}
	)
	type invalid struct {
		Name     int
		Email    string
		Password string
	}
	var (
		data_invalid = invalid{
			Name:     1,
			Email:    "hanief@gmail.com",
			Password: "12345",
		}
	)
	test := testCases{
		path: "/users/:id",
	}

	t.Run("Succes Update User", func(t *testing.T) {

		body, err := json.Marshal(data_user)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPut, test.path, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateUserController(c)) {

			bodyResponses := rec.Body.String()

			var user responsejson.ResponCreate
			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, data_user.Name, user.Data.Name)
			assert.Equal(t, data_user.Email, user.Data.Email)
			assert.Equal(t, "success", user.Status)
		}
	})

	t.Run("Invalid Format", func(t *testing.T) {
		body, err := json.Marshal(data_invalid)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user responsejson.ResponCreate

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})

	t.Run("Invalid Convert", func(t *testing.T) {
		body, err := json.Marshal(data_user)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("aku")

		if assert.NoError(t, UpdateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user responsejson.ResponCreate

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})

	t.Run("Failed Update User", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("80")

		if assert.NoError(t, UpdateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user responsejson.ResponCreate

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})
}

func TestDeleteUserController(t *testing.T) {
	e := InitEchoTestAPINoDop()

	test := testCases{
		path: "/users/:id",
	}

	t.Run("Succes Delete User", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, test.path, nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteUserController(c)) {

			bodyResponses := rec.Body.String()

			var user responsejson.ResponCreate
			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success", user.Status)
		}
	})

	t.Run("Failed Delete User", func(t *testing.T) {

		config.DB.Migrator().DropTable(&model.User{})
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("80")

		if assert.NoError(t, UpdateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user responsejson.ResponCreate

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})
}
