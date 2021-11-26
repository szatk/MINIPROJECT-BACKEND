package config

import (
	"MINIPROJECT-BACKEND/model/houseperca"
	"MINIPROJECT-BACKEND/model/transaksi"
	"MINIPROJECT-BACKEND/model/userlogin"
	"MINIPROJECT-BACKEND/model/users"
)

func Migrate() {
	DB.AutoMigrate(
		&users.User{},
		&userlogin.LoginDataUsers{},
		&houseperca.HousePerca{},
		&houseperca.PekerjaPerca{},
		&transaksi.JenisPerca{},
		&transaksi.Transaksi{},
		&transaksi.DetailTransaksi{},
	)
}