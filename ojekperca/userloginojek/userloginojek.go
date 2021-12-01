package UserLoginDriver

import (
	"MINIPROJECT-BACKEND/model/users"
	"errors"

	"gorm.io/gorm"
)

func CreateUserLogin(db *gorm.DB, u *users.LoginDataUsers) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUserLogins(db *gorm.DB) ([]users.LoginDataUsers, error) {
	userLogins := []users.LoginDataUsers{}

	if err := db.Preload("User").Find(&userLogins).Error; err != nil {
		return userLogins, err
	}

	return userLogins, nil
}

func GetUserLoginByID(id string, db *gorm.DB) (users.LoginDataUsers, bool, error) {
	var userlogin users.LoginDataUsers

	err := db.Preload("User").First(&userlogin, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return userlogin, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userlogin, false, nil
	}
	return userlogin, true, nil
}

func DeleteUserLogin(id string, db *gorm.DB) error {
	var userLogin users.LoginDataUsers
	if err := db.Where("id = ? ", id).Unscoped().Delete(&userLogin).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserLogin(db *gorm.DB, u *users.LoginDataUsers) error {
	if err := db.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func Login(db *gorm.DB, log *users.Login, ul *users.LoginDataUsers) (users.LoginDataUsers, error) {
	if err := db.Where("username = ? OR email = ?", log.Username, log.Email).Find(ul).Error; err != nil {
		return *ul, err
	}
	return *ul, nil
}
