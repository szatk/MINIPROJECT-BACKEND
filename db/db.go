package db

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/model"
)

func CreateUser(user model.UserLogin) model.UserLogin {
	config.DB.Create(&user)
	return user
}
