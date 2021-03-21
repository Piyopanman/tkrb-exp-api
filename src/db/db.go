package db

import (
	"database/sql"
	"fmt"
	"os"
	"tkrb-exp-api/src/logging"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Conn DB接続情報
var Conn *gorm.DB

//Init データベースの初期化
func Init(){
	// err := godotenv.Load()
	// if err != nil {
	// 	logging.Logger.Error("Error loading .env file")
	// }
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
	url := os.Getenv("DATABASE_URL")
	connection,err := pq.ParseURL(url)
	if err != nil{
		logging.Logger.Error("Failed to parse db URL")
		return
	}
	connection += " sslmode=require"
	fmt.Println(connection)
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


	// host := "127.0.0.1"
	// user := "kakizakihinano"
	// pass := "ishikari0719"
	// dbname := "touken_api"
	// port := "5432"
	// var err error
	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user,pass, dbname,port)
	// 	Conn,err = gorm.Open(postgres.Open(dns), &gorm.Config{
	// 	CreateBatchSize: 100,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	if err != nil{
		logging.Logger.Error("Failed to connect to SQL")
		return
	}

	logging.Logger.Info("Success to connect to SQL")
	
}