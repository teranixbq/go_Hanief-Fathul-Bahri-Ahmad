package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"praktikum/rest/database"
	"praktikum/rest/model"
	"praktikum/rest/controller/responsejson"
	"strconv"
)


// create new blogs
func CreateBlogsController(c echo.Context) error {
	blogs := model.Blogs{}
	c.Bind(&blogs)

	err := database.InsertBlogs(blogs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blogs",
	})
}


// get all blogs
func GetAllBlogsController(c echo.Context) error {

	response, err := database.SelectAllBlogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
    var blogsResponse = []responsejson.ResponseBlogs{}
    for _, value := range response {
        blogsResponse = append(blogsResponse, responsejson.ResponseBlogs{
            ID:        value.ID,
            Judul:     value.Judul,
            Konten:    value.Konten,
            // UserID:    value.UserID,
            CreatedAt: value.CreatedAt,
            User: responsejson.ResponseUser{
				ID:        value.User.ID,
				Name:      value.User.Name,
				Email:     value.User.Email,
				CreatedAt: value.User.CreatedAt,
        },
        })
    }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogss":   blogsResponse,
	})
}

// get blogs by id
func GetBlogsByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	response, err := database.SelectBlogById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ada",
		})
	}

	var blogsResponse = []responsejson.ResponseBlogs{}
	for _, value := range response{
		blogsResponse = append (blogsResponse, responsejson.ResponseBlogs{
			ID:        value.ID,
			Judul:     value.Judul,
			Konten:    value.Konten,
			CreatedAt: value.CreatedAt,
			User: responsejson.ResponseUser{
				ID:        value.User.ID,
				Name:      value.User.Name,
				Email:     value.User.Email,
				CreatedAt: value.User.CreatedAt,
			},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blogss":  blogsResponse,
	})
}


// update blogs by id
func UpdateBlogsController(c echo.Context) error {
	updateData := model.Blogs{}
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

	err = database.UpdateBlogs(id, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})

}


// delete blogs by id
func DeleteBlogsByIdController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}
	err = database.DeleteBlogs(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
