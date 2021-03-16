package handler

import (
	"log"
	"math"
	"net/http"
	"touken-exp/src/constant"
	"touken-exp/src/model"

	"github.com/gin-gonic/gin"
)

//GetToukenOne 一振り
func GetToukenOne(c *gin.Context){
	//リクエストボディから値を取得
	var req getToukenOneRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		log.Println("failed to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message":"failed!!!!!!"})
		return
	}

	//刀剣IDから刀剣名、刀種IDを取得
	var touken model.Touken
	touken,err := model.GetTouken(req.Touken)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message":"internal server error"})
		return
	}

	//刀種IDとレベルから経験値を取得
	var exp model.Exp
	exp,err = model.GetExp(touken.ToushuID,req.Level)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message":"internal server error"})
		return
	}

	//金平糖何個分？
	konpeto := float64(exp.SumExp) / constant.KonpetoExp

	//厚樫山何周分？
	atsukashiKari := float64(exp.SumExp) / constant.AtsukashiExp
	atsukashi := math.Round(atsukashiKari * 100) / 100


	c.JSON(http.StatusOK, getToukenOneResponse{
		ToukenName: touken.Touken,
		Level: req.Level,
		Exp: exp.SumExp,
		Konpeto: konpeto,
		Atsukashi: atsukashi,
	})


}

type getToukenOneRequest struct {
	Saniwa string `json:"saniwa"`
	Touken int `json:"touken,string,omitempty"`
	Level int `json:"level,string,omitempty"`
}

type getToukenOneResponse struct{
	ToukenName string `json:"toukenName"`
	Level int `json:"level"`
	Exp int32 `json:"exp"`
	Konpeto float64 `json:"konpeto"`
	Atsukashi float64 `json:"atsukashi"`
}