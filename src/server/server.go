package server

import (
	"touken-exp/src/handler"
	"touken-exp/src/middleware"

	"github.com/gin-gonic/gin"
)

//Init サーバー起動
func Init(){
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine{
	router := gin.Default()
	router.Use(middleware.Cors())

	//* ルーティング *//
	// router.GET("/",handler.Top) //ほんとはいらない
	// router.GET("/hitokuchi", handler.Hito)
	// router.POST("/futakuchi",handler.Futa)

	router.GET("/toukenList", handler.GetToukenList)
	router.POST("/toukenOne",handler.GetToukenOne)
	router.POST("/toukenTwo", handler.GetToukenTwo)

	return router
}