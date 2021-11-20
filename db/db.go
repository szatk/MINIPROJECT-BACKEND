package db

import (
	"MINIPROJECT-BACKEND/config"
	"MINIPROJECT-BACKEND/model"
)

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}