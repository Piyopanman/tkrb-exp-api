package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetToukenOne 一振り
func GetToukenOne(c *gin.Context){
	fmt.Println("getToukenOne is called")
	//リクエストボディから値を取得
	var req getToukenOneRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		log.Println("failed to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message":"failed!!!!!!"})
		return
	}

	fmt.Println(req.Saniwa)
	fmt.Println(req.Touken)
	fmt.Println(req.Level)
	//ここのデータベース処理を書くところから始める（返すのは、刀剣名、経験値、金平糖、厚樫山、...）
	c.JSON(http.StatusOK, gin.H{"message":"ok!"})

}

type getToukenOneRequest struct {
	Saniwa string `json:"saniwa"`
	Touken int `json:"touken,string,omitempty"`
	Level int `json:"level,string,omitempty"`
}

type getToukenOneResponse struct{

}