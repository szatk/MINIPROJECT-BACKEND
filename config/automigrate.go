package config

import (
	"MINIPROJECT-BACKEND/model/houseperca"
	"MINIPROJECT-BACKEND/model/transaksi"
	"MINIPROJECT-BACKEND/model/users"
)

// fungsi bikin tabel otomatis
func InitMigration() {
	DB.AutoMigrate(
		&users.User{},
		&users.LoginDataUsers{},
		&houseperca.HousePerca{},
		&houseperca.PekerjaPerca{},
		&transaksi.JenisPerca{},
		&transaksi.Transaksi{},
		&transaksi.DetailTransaksi{},
	)
}
