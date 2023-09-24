package main

import (
  "net/http"
  "strconv"
  "github.com/labstack/echo"
)

type User struct {
  Id    int    	  `json:"id" 	   form:"id"`
  Name  string 	  `json:"name" 	   form:"name"`
  Email string 	  `json:"email"    form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUsersControllerID(c echo.Context) error {
  // your solution here
  var userById User
  id,err := strconv.Atoi(c.Param("id"))
  if err != nil {
    // Jika konversi gagal, mengembalikan respons kesalahan
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "Failed Convert",
    })
   }
   
   found := false
   for _,value:= range users{
		if value.Id == id{
			userById = value
			found = true
			break
		}
   }

   if found {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes",
			"user" : userById,
		})
   }else {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Tidak Ada",
		})
   }


}
// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  id,err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "Failed Convert",
    })
   }

   for i, user := range users{
		if user.Id == id {
			users = append(users[:i],users[i+1:]... )
			
			return c.JSON(http.StatusOK, map[string]interface{}{
      			"message": "Delete Succes",
    		})
		}
   }

   return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data Tidak Ada",
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "Failed convert",
    })
  }
  
  found := false
  for i, value := range users {
    if value.Id == id {
      if err := c.Bind(&value); err != nil {

        return c.JSON(http.StatusBadRequest, map[string]interface{}{
          "message": "Invalid format",
        })
      }

      users[i] = value
      found = true
      break
    }
  }

  if !found {
    return c.JSON(http.StatusNotFound, map[string]interface{}{
      "message": "User tidak ada",
    })
  }

  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success update data",
  })

}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  if err := c.Bind(&user); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "Invalid format",
    })
  }

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}

// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.GET("/users/:id", GetUsersControllerID)
  e.POST("/users", CreateUserController)
  e.PUT("/users/:id", UpdateUserController)
  e.DELETE("/users/:id", DeleteUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}