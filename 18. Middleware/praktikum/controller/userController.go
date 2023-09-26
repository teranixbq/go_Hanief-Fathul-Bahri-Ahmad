package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum/rest/database"
	"praktikum/rest/model"
	"praktikum/rest/controller/responsejson"
	"praktikum/rest/middleware"
	"strconv"
)

func LoginUserController(c echo.Context) error {
    user := model.User{}
    c.Bind(&user)

    data, token, err := database.UserLoginAuth(user.Email, user.Password)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "message": "Login failed: Invalid email or password",
        })
    }

	response := map[string]any{
			"user_id": data.ID,
			"name":    data.Name,
			"token":   token,
	}


    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Login successful",
        "user":    response,  
    })
}


// get all users
func GetUsersController(c echo.Context) error {

	response, err := database.SelectAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	var userResponse = []responsejson.ResponseUser{}
	for _, value := range response {
		userResponse = append(userResponse,responsejson.ResponseUser{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   userResponse,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userID, name := middleware.ExtractToken(c)

    if userID == 0 || name == "" {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	if userID != id {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": "Access denied",
        })
    }

	response, err := database.SelectById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data tidak ada",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"users":   response,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := database.InsertUsers(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userID, name := middleware.ExtractToken(c)

    if userID == 0 || name == "" {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	if userID != id {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": "Access denied",
        })
    }

	err = database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {


	userID, name := middleware.ExtractToken(c)

    if userID == 0 || name == "" {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized",
        })
    }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Convert",
		})
	}

	if userID != id {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": "Access denied",
        })
    }
	
	updateData := model.User{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid Format",
		})
	}

	err = database.UpdateUser(id, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
	})

}
