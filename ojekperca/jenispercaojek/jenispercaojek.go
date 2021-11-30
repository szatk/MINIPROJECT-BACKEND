package jenispercaojek

import (
	"MINIPROJECT-BACKEND/model/transaksi"
	"errors"

	"gorm.io/gorm"
)

func CreateJenisPerca(db *gorm.DB, js *transaksi.JenisPerca) error {
	if err := db.Create(&js).Error; err != nil {
		return err
	}
	return nil
}

func GetAllJenisPerca(db *gorm.DB) ([]transaksi.JenisPerca, error) {
	jenisPerca := []transaksi.JenisPerca{}

	if err := db.Find(&jenisPerca).Error; err != nil {
		return jenisPerca, err
	}

	return jenisPerca, nil
}

func GetJenisPercaByID(id string, db *gorm.DB) (transaksi.JenisPerca, bool, error) {
	js := transaksi.JenisPerca{}

	err := db.Where("id = ?", id).First(&js).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return js, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return js, false, nil
	}
	return js, true, nil
}

func UpdateJenisPerca(db *gorm.DB, js *transaksi.JenisPerca, id string) error {
	if err := db.Model(transaksi.JenisPerca{}).Where("id = ?", id).Updates(js).Error; err != nil {
		return err
	}
	return nil
}

func DeleteJenisPerca(id string, db *gorm.DB) error {
	var js transaksi.JenisPerca
	if err := db.Where("id = ? ", id).Unscoped().Delete(&js).Error; err != nil {
		return err
	}
	return nil
}
