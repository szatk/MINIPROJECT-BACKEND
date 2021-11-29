package config

import (
	"MINIPROJECT-BACKEND/model/houseperca"


// fungsi bikin tabel otomatis
func InitMigration() {
	DB.AutoMigrate(
		&houseperca.HousePerca{},
		&houseperca.PekerjaPerca{},
	)
}
