package controller

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/help"
	"MINIPROJECT-BACKEND/middleware"
	"MINIPROJECT-BACKEND/model/respon"
	"MINIPROJECT-BACKEND/model/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUserLogin ... Post new meta Login
func CreateUserLogin(c echo.Context) error {
	var userlogin users.LoginDataUsers
	var user users.User

	c.Bind(&userlogin)
	if res := config.DB.Where("id = ?", userlogin.UserId).Find(&user); res.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	userlogin.Password = help.Encript(userlogin.Password)
	userlogin.User = user

	if res := config.DB.Create(&userlogin); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Failed create data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Success create userlogin",
		Data:    &userlogin,
	})

}

// GetAllUserVerification ... All User Verifation Data
func GetAllUserLogin(c echo.Context) error {
	var userlogins = []users.LoginDataUsers{}

	if res := config.DB.Joins("User").Find(&userlogins); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful retrieve data",
		Data:    &userlogins,
	})
}

// UpdateUserLogin ... Update User Login Data
func GetUserLoginByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	var userlogin users.LoginDataUsers
	result := config.DB.First(&userlogin, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &userlogin,
	})
}

// UserLoginValidation ... All User Verifation Data
func Login(c echo.Context) error {
	var login users.Login
	c.Bind(&login)

	// check form input
	if login.Email == "" && login.Username == "" {
		return c.JSON(http.StatusBadRequest, respon.BaseRespon{
			Code:    http.StatusBadRequest,
			Message: "invalid input",
			Data:    nil,
		})
	} else if login.Password == "" {
		return c.JSON(http.StatusBadRequest, respon.BaseRespon{
			Code:    http.StatusBadRequest,
			Message: "invalid input",
			Data:    nil,
		})
	}

	// check data to database
	var userlogin users.LoginDataUsers
	res := config.DB.Where("username = ? OR email = ?", login.Username, login.Email).Find(&userlogin)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	if userlogin.Email == "" || userlogin.Username == "" || userlogin.Password == "" {
		return c.JSON(http.StatusForbidden, respon.BaseRespon{
			Code:    http.StatusForbidden,
			Message: "access denied: email or username invalid",
			Data:    nil,
		})
	}

	// compare password with hashed
	access := help.ComparePassword(userlogin.Password, login.Password)
	if !access {
		return c.JSON(http.StatusForbidden, respon.BaseRespon{
			Code:    http.StatusForbidden,
			Message: "access denied: password did not match",
			Data:    nil,
		})
	}

	// get token
	tokenLogin, err := middleware.CreateToken(int(userlogin.UserId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "failed create token",
			Data:    nil,
		})
	}

	// set login response
	loginResponse := respon.LoginRespon{
		Id:        int(userlogin.UserId),
		Username:  userlogin.Username,
		Email:     userlogin.Email,
		Token:     tokenLogin,
		CreatedAt: userlogin.CreatedAt,
		UpdatedAt: userlogin.UpdatedAt,
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "access complete",
		Data:    &loginResponse,
	})
}

// UpdateUserLogin ... Update User Login Data
func UpdateUserLogin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	var userlogin users.LoginDataUsers
	result := config.DB.First(&userlogin, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&userlogin)
	userlogin.Password = help.Encript(userlogin.Password)

	result = config.DB.Save(&userlogin)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &userlogin,
	})
}
