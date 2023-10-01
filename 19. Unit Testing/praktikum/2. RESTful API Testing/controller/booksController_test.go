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

type testBooks struct {
	path       string
	expectCode int
}

func initEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func initEchoTestAPINoDop() *echo.Echo {
	config.InitDBTestNoDrop()
	e := echo.New()
	return e
}

var (
	data_books = model.Books{
		Judul:    "alta",
		Penulis:  "alta@gmail.com",
		Penerbit: "12345",
	}
)

func InsertMockDataBooks() error {
	var err error
	if err = config.DB.Create(&data_books).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateBooksController(t *testing.T) {
	test := testBooks{
		path: "/books",
	}
	t.Run("Success Create", func(t *testing.T) {

		test.expectCode = http.StatusOK

		e := initEchoTestAPI()
		body, err := json.Marshal(data_books)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPost, test.path, bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBooksController(c)) {

			bodyResponses := rec.Body.String()

			var books responsejson.ResponCreateBooks
			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.expectCode, rec.Code)
			assert.Equal(t, data_books.Judul, books.Data.Judul)
			assert.Equal(t, data_books.Penulis, books.Data.Penulis)
			assert.Equal(t, data_books.Penerbit, books.Data.Penerbit)
			assert.Equal(t, "success", books.Status)
		}
	})

	t.Run("failed Create", func(t *testing.T) {

		test.expectCode = http.StatusBadRequest

		e := InitEchoTestAPINoDop()
		config.DB.Migrator().DropTable(&model.Books{})

		req := httptest.NewRequest(http.MethodPost, test.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call function on controller
		if assert.NoError(t, CreateBooksController(c)) {
			bodyResponses := rec.Body.String()
			var books responsejson.ResponCreateBooks

			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, test.expectCode, rec.Code)
			
		}
	})

}

func TestGetAllBooksController(t *testing.T) {
	test := testBooks{
		path: "/books",
	}
	t.Run("Success Get All Books", func(t *testing.T) {

		test.expectCode = http.StatusOK

		e := initEchoTestAPI()
		InsertMockDataBooks()
		req := httptest.NewRequest(http.MethodGet, test.path, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetAllBooksController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}

	})

	t.Run("Failed Get All Books", func(t *testing.T) {
		test.expectCode = http.StatusInternalServerError

		e := initEchoTestAPI()
		config.DB.Migrator().DropTable(&model.Books{})

		req := httptest.NewRequest(http.MethodPost, test.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call function on controller
		if assert.NoError(t, GetAllBooksController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}

	})
}

func TestGetBooksByIdController(t *testing.T) {
	e := initEchoTestAPI()
	test := testBooks{
		path: "/books/:id",
	}

	t.Run("Succes GetBooksById", func(t *testing.T) {
		test.expectCode = http.StatusOK

		e := initEchoTestAPI()
		InsertMockDataBooks()
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(test.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetBooksByIdController(c)) {

			assert.Equal(t, test.expectCode, rec.Code)
		}
	})

	t.Run("Failed GetBooksById", func(t *testing.T) {
		test.expectCode = http.StatusBadRequest

		InsertMockDataBooks()
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("#")

		if assert.NoError(t, GetBooksByIdController(c)) {

			assert.Equal(t, test.expectCode, res.Code)

		}
	})

}

func TestUpdateBooksController(t *testing.T) {
	e := initEchoTestAPINoDop()

	var (
		data_books = model.Books{
			Judul:    "Ekonomi",
			Penulis:  "Eko",
			Penerbit: "DB",
		}
	)
	test := testBooks{
		path: "/books/:id",
	}

	t.Run("Succes Update Books", func(t *testing.T) {

		body, err := json.Marshal(data_books)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPut, test.path, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateBooksController(c)) {

			bodyResponses := rec.Body.String()

			var books responsejson.ResponCreateBooks
			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, data_books.Judul, books.Data.Judul)
			assert.Equal(t, data_books.Penulis, books.Data.Penulis)
			assert.Equal(t, data_books.Penerbit, books.Data.Penerbit)
			assert.Equal(t, "success", books.Status)
		}
	})

	t.Run("Failed Update Books", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("80")

		if assert.NoError(t, UpdateBooksController(c)) {
			bodyResponses := rec.Body.String()
			var books responsejson.ResponCreateBooks

			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})
}

func TestDeleteBooksController(t *testing.T) {
	e := initEchoTestAPINoDop()

	test := testBooks{
		path: "/books/:id",
	}

	t.Run("Succes Delete Books", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, test.path, nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteBooksByIdController(c)) {

			bodyResponses := rec.Body.String()

			var books responsejson.ResponCreateBooks
			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success", books.Status)
		}
	})

	t.Run("Failed Delete Books", func(t *testing.T) {

		config.DB.Migrator().DropTable(&model.Books{})
		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("80")

		if assert.NoError(t, UpdateBooksController(c)) {
			bodyResponses := rec.Body.String()
			var books responsejson.ResponCreateBooks

			err := json.Unmarshal([]byte(bodyResponses), &books)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}
	})
}
