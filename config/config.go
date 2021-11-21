package config

import (
	"MINIPROJECT-BACKEND/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// fungsi connect ke mySQL
func InitDB() {
	dsn := "root:567890@tcp(127.0.0.1:3306)/atikahperca?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// fungsi bikin tabel otomatis
func InitMigration() {
	DB.AutoMigrate(
		&model.UserLogin{},
		&model.User{},
	)
}

// type Configuration struct {
// 	DB_USERNAME string
// 	DB_PASSWORD string
// 	DB_PORT     string
// 	DB_HOST     string
// 	DB_NAME     string
// }

// func GetConfig() Configuration {
// 	configuration := Configuration{
// 		DB_NAME: 'test_go',
// 		DB_USERNAME: 'route',
// 		DB_PORT: '3306',
// 		DB_HOST: 'localhost',
// 		DB_PASSWORD: '',
// 	}
// 	return configuration
// }
