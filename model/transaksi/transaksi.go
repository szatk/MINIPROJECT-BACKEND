package transaksi

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	Id              uint           `gorm:"PrimaryKey; Auto Increment; Not Null" json:"id"`
	HousePercaId    uint           `gorm:"not null" json:"housePercaId"`
	UserId          uint           `gorm:"not null; ForeignKey; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"userId"`
	PekerjaPercaId  uint           `gorm:"not null" json:"pekerjaId"`
	TotalQty        int            `gorm:"size:10" json:"totalQty"`
	TotalTransaksi  float32        `gorm:"type:decimal(13,2)" json:"totalTransaksi"`
	Status          string         `gorm:"size:255;default:process" json:"status"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DetailTransaksi []DetailTransaksi
}

type DetailTransaksi struct {
	TransaksiID  uint    `gorm:"Not Null; PrimaryKey; autoIncrement:false; ForeignKey; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transactionId"`
	JenisPercaId uint    `gorm:"Not Null; PrimaryKey; autoIncrement:false" json:"jenisPercaId"`
	Qty          int     `gorm:"size:10; not null" json:"qty"`
	TotalPrice   float32 `gorm:"type:decimal(13,2)" json:"totalPrice"`
}

type TransaksiReq struct {
	HousePercaId   uint              `json:"housePercaId"`
	UserId         uint              `json:"userId"`
	PekerjaPercaId uint              `json:"pekerjaId"`
	TotalQty       int               `json:"totalQty"`
	TotalTransaksi float32           `json:"totalTransaksi"`
	Status         string            `json:"status"`
	Detail         []DetailTransaksi `json:"detail"`
}
