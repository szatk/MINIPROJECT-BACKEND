package housepercaojek

import (
	"MINIPROJECT-BACKEND/model/houseperca"
	"errors"

	"gorm.io/gorm"
)

func CreateHousePerca(db *gorm.DB, b *houseperca.HousePerca) error {
	if err := db.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllHousePerca(db *gorm.DB) ([]houseperca.HousePerca, error) {
	housePerca := []houseperca.HousePerca{}

	if err := db.Find(&housePerca).Error; err != nil {
		return housePerca, err
	}

	return housePerca, nil
}

func GetHousePercaByID(id string, db *gorm.DB) (houseperca.HousePerca, bool, error) {
	b := houseperca.HousePerca{}

	err := db.Where("id = ?", id).First(&b).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteHousePerca(id string, db *gorm.DB) error {
	var b houseperca.HousePerca
	if err := db.Where("id = ? ", id).Unscoped().Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateHousePerca(db *gorm.DB, b *houseperca.HousePerca, id string) error {
	if err := db.Model(houseperca.HousePerca{}).Where("id = ?", id).Updates(b).Error; err != nil {
		return err
	}
	return nil
}
