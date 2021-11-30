package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBLog *mongo.Client

// fungsi connect ke mysql
func InitDB() {
	dsn := "root:567890@tcp(127.0.0.1:3306)/atikahperca?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

//fungsi connect mongodb
func InitLog() {
	var err error
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/atikahperca"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
}
