package operatorpercadriver

import (
	"MINIPROJECT-BACKEND/model/houseperca"
	"errors"

	"gorm.io/gorm"
)

// Register Operator Perca
func CreateOperatorPerca(db *gorm.DB, os *houseperca.OperatorPerca) error {
	if err := db.Create(&os).Error; err != nil {
		return err
	}
	return nil
}

//Menampilkan All Operator Perca
func GetAllOperatorPerca(db *gorm.DB) ([]houseperca.OperatorPerca, error) {
	operatorPerca := []houseperca.OperatorPerca{}

	if err := db.Preload("houseperca").Find(&operatorPerca).Error; err != nil {
		return operatorPerca, err
	}

	return operatorPerca, nil
}

//Menampilkan Operator Perca by ID
func GetOperatorPercaByID(id string, db *gorm.DB) (houseperca.OperatorPerca, bool, error) {
	var operatorPerca houseperca.OperatorPerca

	err := db.Preload("houseperca").First(&operatorPerca, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return operatorPerca, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operatorPerca, false, nil
	}
	return operatorPerca, true, nil
}

//Update Operator Perca
func UpdateOperatorPerca(id string, db *gorm.DB, op *houseperca.OperatorPerca) error {
	if err := db.Model(houseperca.OperatorPerca{}).Where("id = ?", id).Updates(op).Error; err != nil {
		return err
	}
	return nil
}

//Delete Operator Perca
func DeleteOperatorPerca(id string, db *gorm.DB) error {
	var operatorSampah houseperca.OperatorPerca
	if err := db.Where("id = ? ", id).Unscoped().Delete(&operatorSampah).Error; err != nil {
		return err
	}
	return nil
}

//Menampilkan Operator Perca by ID
func GetOperatorPercaByObject(db *gorm.DB, op *houseperca.OperatorPerca) (houseperca.OperatorPerca, error) {
	if err := db.Preload("houseperca").Find(&op).Error; err != nil {
		return *op, err
	}

	return *op, nil
}
