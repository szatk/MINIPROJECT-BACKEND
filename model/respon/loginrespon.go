package respon

import (
	"time"

	"gorm.io/gorm"
)

type LoginRespon struct {
	Id        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Token     string         `json:"token"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
