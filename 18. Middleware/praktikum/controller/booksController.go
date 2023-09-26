package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum/rest/database"
	"praktikum/rest/model"
	"praktikum/rest/controller/responsejson"
	"strconv"
)

// get all books
func GetAllBooksController(c echo.Context) error {

	response, err := database.SelectAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	var bookResponse = []responsejson.ResponseBook{}
	for _, value := range response {
		bookResponse = append(bookResponse, responsejson.ResponseBook{
			ID:        value.ID,
			Judul:     value.Judul,
			Penerbit:  value.Penerbit,
			Penulis:   value.Penulis,
			CreatedAt: value.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   bookResponse,
	})
}

// get book by id
func GetBooksByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	response, err := database.SelectBooksById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ada",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"books":   response,
	})
}

// create new book
func CreateBooksController(c echo.Context) error {
	books := model.Books{}
	c.Bind(&books)

	err := database.InsertBooks(books)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
	})
}

// delete book by id
func DeleteBooksByIdController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}
	err = database.DeleteBooks(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBooksController(c echo.Context) error {
	updateData := model.Books{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Format",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	err = database.UpdateBooks(id, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})

}
