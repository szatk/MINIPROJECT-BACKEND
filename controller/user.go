package controller

import (
	"net/http"
	"MINIPROJECT-BACKEND/db"
	"MINIPROJECT-BACKEND/model"
	"github.com/labstack/echo"
)

// menambah user baru
func CreateUserController(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Can't Create User, Try Again!",
			"error":   err.Error(),
		})
	}

	newUser = db.CreateUser(newUser)
	newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success Creates User",
		"data":    newUser,
	})
}