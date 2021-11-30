package userojek

import (
	"MINIPROJECT-BACKEND/model/users"
	"errors"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, u *users.User) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *gorm.DB) ([]users.User, error) {
	users := []users.User{}

	if err := db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByID(id string, db *gorm.DB) (users.User, bool, error) {
	u := users.User{}

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return u, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, false, nil
	}
	return u, true, nil
}

func DeleteUser(id string, db *gorm.DB) error {
	var u users.User
	if err := db.Where("id = ? ", id).Unscoped().Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, u *users.User, id string) error {
	if err := db.Model(users.User{}).Where("id = ?", id).Updates(u).Error; err != nil {
		return err
	}
	return nil
}
