package pekerjapercaojek

import (
	"MINIPROJECT-BACKEND/model/houseperca"
	"errors"

	"gorm.io/gorm"
)

// Register Pekerja Perca
func CreatePekerjaPerca(db *gorm.DB, os *houseperca.PekerjaPerca) error {
	if err := db.Create(&os).Error; err != nil {
		return err
	}
	return nil
}

func GetAllPekerjaPerca(db *gorm.DB) ([]houseperca.PekerjaPerca, error) {
	pekerjaPerca := []houseperca.PekerjaPerca{}

	if err := db.Preload("BankSampah").Find(&pekerjaPerca).Error; err != nil {
		return pekerjaPerca, err
	}

	return pekerjaPerca, nil
}

func GetPekerjaPercaByID(id string, db *gorm.DB) (houseperca.PekerjaPerca, bool, error) {
	var pekerjaPerca houseperca.PekerjaPerca

	err := db.Preload("HousePerca").First(&pekerjaPerca, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return pekerjaPerca, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pekerjaPerca, false, nil
	}
	return pekerjaPerca, true, nil
}

func UpdatePekerjaPerca(id string, db *gorm.DB, op *houseperca.PekerjaPerca) error {
	if err := db.Model(houseperca.PekerjaPerca{}).Where("id = ?", id).Updates(op).Error; err != nil {
		return err
	}
	return nil
}

func DeletePekerjaPerca(id string, db *gorm.DB) error {
	var pekerjaPerca houseperca.PekerjaPerca
	if err := db.Where("id = ? ", id).Unscoped().Delete(&pekerjaPerca).Error; err != nil {
		return err
	}
	return nil
}

func GetPekerjaPercaByObject(db *gorm.DB, op *houseperca.PekerjaPerca) (houseperca.PekerjaPerca, error) {
	if err := db.Preload("HousePerca").Find(&op).Error; err != nil {
		return *op, err
	}

	return *op, nil
}
