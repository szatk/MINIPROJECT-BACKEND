package model

import (
	"time"
	//"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	ID        uint      `gorm:"primarykey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Token     string    `json:"token,"`
}
