package db

import (
	"database/sql"
	"os"
	"tkrb-exp-api/src/logging"

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
		logging.Logger.Info("product mode DB")
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
		logging.Logger.Info("Success to connect to SQL")
	
}