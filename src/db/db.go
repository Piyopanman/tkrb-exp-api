package db

import (
	"database/sql"
	"fmt"
	"os"
	"tkrb-exp-api/src/logging"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Conn DB接続情報
var Conn *gorm.DB

//Init データベースの初期化
func Init(){

	//postgreSQL
	url := os.Getenv("DATABASE_URL")
	if(url != ""){
		sqlDB,err := sql.Open("postgres", url)
		Conn, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		  }), &gorm.Config{
			CreateBatchSize: 100,
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
		  })
		if err != nil{
			logging.Logger.Error("Failed to open database")
			return
		}
	}else{
		err := godotenv.Load()
		if err != nil {
			logging.Logger.Error("Error loading .env file")
			return
		}
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user,pass, dbname,port)
			Conn,err = gorm.Open(postgres.Open(dns), &gorm.Config{
			CreateBatchSize: 100,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil{
			logging.Logger.Error("Failed to connect to SQL")
			return
		}	
	}
	logging.Logger.Info("Success to connect to SQL")
	
}