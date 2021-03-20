package server

import (
	"os"
	"tkrb-exp-api/src/handler"
	"tkrb-exp-api/src/middleware"

	"github.com/gin-gonic/gin"
)

//Init サーバー起動
func Init(){
	r := router()
	port := os.Getenv("PORT")
	if port == ""{
		port = "5000"
	}
	r.Run(":" + port)
}

func router() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.Cors(), middleware.AccessLog())

	//* ルーティング *//
	router.GET("/toukenList", handler.GetToukenList)
	router.POST("/toukenOne",handler.GetToukenOne)
	router.POST("/toukenTwo", handler.GetToukenTwo)

	return router
}