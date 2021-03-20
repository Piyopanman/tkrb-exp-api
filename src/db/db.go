package db

import (
	"fmt"
	"touken-exp/src/logging"

	_ "github.com/go-sql-driver/mysql" //blank import
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Conn DB接続情報
var Conn *gorm.DB

//Init データベースの初期化
func Init(){
	err := godotenv.Load()
	if err != nil {
		logging.Logger.Error("Error loading .env file")
	}
	//mysql
	// user := os.Getenv("MYSQL_USER")
	// password := os.Getenv("MYSQL_PASSWORD")
	// // host := os.Getenv("MYSQL_HOST")
	// host := "mysql"
	// port := os.Getenv("MYSQL_PORT")
	// database := os.Getenv("MYSQL_DATABASE")

	// dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	// Conn,err = gorm.Open(mysql.Open(dns), &gorm.Config{
	// 	CreateBatchSize: 100,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	//postgreSQL
	dns := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tokyo")
	Conn,err =  gorm.Open(postgres.Open(dns), &gorm.Config{
			CreateBatchSize: 100,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

	if err != nil{
		logging.Logger.Error("Failed to connect to SQL")
	}

	logging.Logger.Info("Success to connect to SQL")
	
}