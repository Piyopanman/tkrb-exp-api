package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //blank import
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Conn DB接続情報
var Conn *gorm.DB

//Init データベースの初期化
func Init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	// host := os.Getenv("MYSQL_HOST")
	host := "mysql"
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	Conn,err = gorm.Open(mysql.Open(dns), &gorm.Config{
		CreateBatchSize: 100,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil{
		log.Println(err)
		log.Println("failed to connect to SQL")
	}
	
}