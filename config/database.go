package config

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

import (

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var DBLog *mongo.Client

// fungsi connect ke mysql
func InitDB() {
	dsn := "root:567890@tcp(127.0.0.1:3306)/atikahperca?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// //fungsi connect mongodb
// func InitLog() {
// 	var err error
// 	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/study-timer"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
// 	err = DBLog.Connect(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	DBLog.ListDatabaseNames(ctx, bson.M{})
// }

// var DB *gorm.DB
// var err error

// // DBConfig represents db configuration
// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

// func BuildDBConfig() *DBConfig {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	host := os.Getenv("DB_HOST")
// 	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")

// 	dbConfig := DBConfig{
// 		Host:     host,
// 		Port:     port,
// 		User:     user,
// 		Password: password,
// 		DBName:   dbName,
// 	}
// 	return &dbConfig
// }
// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }

// func Connection() {
// 	dsn := DbURL(BuildDBConfig())
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println("Status:", err)
// 	}
// 	Migrate()
// }
