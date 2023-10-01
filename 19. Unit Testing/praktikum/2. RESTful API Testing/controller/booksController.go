package controller

import (
	"net/http"
	"praktikum/rest/controller/responsejson"
	"praktikum/rest/repository"
	"praktikum/rest/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetAllBooksController(c echo.Context) error {

	response, err := repository.SelectAllBooks()
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

	response, err := repository.SelectBooksById(id)
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

	response := responsejson.ResponCreateBooks{}
	err := repository.InsertBooks(books)
	if err != nil {
		response.Status = "failed"
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	response.Status = "success"
	response.Data = books
	return c.JSON(http.StatusOK, response)
}

// delete book by id
func DeleteBooksByIdController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}
	response := responsejson.ResponCreateBooks{}
	err = repository.DeleteBooks(id)
	if err != nil {
		response.Status = "failed"
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response.Status = "success"
	return c.JSON(http.StatusOK, response)

}

// update book by id
func UpdateBooksController(c echo.Context) error {
	updateData := model.Books{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Format",
		})
	}
	response := responsejson.ResponCreateBooks{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Status = "failed"
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	err = repository.UpdateBooks(id, updateData)
	if err != nil {
		response.Status = "failed"
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	response.Status = "success"
	response.Data = updateData
	return c.JSON(http.StatusOK, response)

}
