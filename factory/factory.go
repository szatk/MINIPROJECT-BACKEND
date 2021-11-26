package factory

import (
	"MINIPROJECT-BACKEND/features/driver/bussiness"
	"MINIPROJECT-BACKEND/features/driver/data"
	"MINIPROJECT-BACKEND/features/driver/presentation"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Presentation struct {
	DriverPresenter presentation.PekerjaPercaPresentasi
}

var DB *gorm.DB

func InitDB() {
	dsn := "root:567890@tcp(127.0.0.1:3306)/atikahperca?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(
		&data.PekerjaPerca{},
	)
}

func Init() Presentation {
	InitDB()
	driverData := data.NewPekerjaPercaData(DB)
	driverBisnis := bussiness.NewPekerjaPercaBisnis(driverData)
	// driverPresentasi := presentation.NewPekerjaPercaPresentasi(driverBisnis)
	return Presentation{
		DriverPresenter: *presentation.NewPekerjaPercaPresentasi(driverBisnis),
	}
}
