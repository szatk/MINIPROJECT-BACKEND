package model

import (
	"gorm.io/gorm"
	//"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	UserLoginID    uint
	NIK            string `gorm:"size:255; not null; unique" json:"nik"`
	NamaDepan      string `gorm:"size:255; not null" json:"namaDepan"`
	NamaBelakang   string `gorm:"size:255; not null" json:"namaBelakang"`
	TanggalLahir   string `gorm:"type:date; not null" json:"tanggalLahir"`
	NoTelepon      string `gorm:"size:255; not null" json:"noTelepon"`
	Alamat         string `gorm:"size:1000; not null" json:"alamat"`
	Kabupaten_Kota string `gorm:"size:255; not null" json:"kabupaten_kota"`
	Provinsi       string `gorm:"size:255; not null" json:"provinsi"`
}
