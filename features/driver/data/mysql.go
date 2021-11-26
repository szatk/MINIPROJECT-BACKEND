package data

import (
	"gorm.io/gorm"
	"MINIPROJECT-BACKEND/features/driver"
)


type PekerjaPercaData struct {
	connection *gorm.DB
}

func NewPekerjaPercaData (db *gorm.DB) driver.Data {
	return &PekerjaPercaData{db}
}

func (percadata *PekerjaPercaData) CreatePekerjaPerca(data driver.PekerjaPerca) error {
	recorddata:= fromCore(data)
	if err := percadata.connection.Create(&recorddata).Error; err != nil {
		return err
	}
	return nil
}

