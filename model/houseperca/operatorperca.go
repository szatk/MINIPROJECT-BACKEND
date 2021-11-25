package houseperca

import (
	"time"

	"gorm.io/gorm"
)

type OperatorPerca struct {
	Id             uint           `gorm:"primaryKey; unique; not null" json:"id"`
	NIK            string         `gorm:"size:255; not null; unique" json:"nik"`
	HousePercaId   uint           `json:"housePercaId"`
	NamaDepan      string         `gorm:"size:255; not null" json:"namaDepan"`
	NamaBelakang   string         `gorm:"size:255;not null" json:"namaBelakang"`
	TanggalLahir   string         `gorm:"type:date;not null" json:"tanggalLahir"`
	NoTelepon      string         `gorm:"size:255;not null; unique" json:"noTelepon"`
	Alamat         string         `gorm:"size:1000;not null" json:"alamat"`
	Kecamatan      string         `gorm:"size:255; not null" json:"kecamatan"`
	Kabupaten_Kota string         `gorm:"size:255;not null" json:"kabupaten_kota"`
	Provinsi       string         `gorm:"size:255;not null" json:"provinsi"`
	HousePerca     HousePerca     `gorm:"foreignKey:HousePercaId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
