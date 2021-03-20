package handler

import (
	"fmt"
	"math"
	"net/http"
	"tkrb-exp-api/src/constant"
	"tkrb-exp-api/src/logging"
	"tkrb-exp-api/src/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//GetToukenOne 一振り
func GetToukenOne(c *gin.Context){
	logging.Logger.Info(fmt.Sprintf("Get access to %v", c.Request.URL.Path))
	//リクエストボディから値を取得
	var req getToukenOneRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		logging.Logger.Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"message":"Failed to bind JSON"})
		return
	}

	//刀剣IDから刀剣名、刀種IDを取得
	var touken model.Touken
	touken,err := model.GetTouken(req.Touken)
	if err != nil{
		logging.Logger.Error(fmt.Sprintf("Failed to get toukenName and toukenID by toukenID(%d)", req.Touken))
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to get toukenName and toukenID in GetToukenOne"})
		return
	}

	//刀種IDとレベルから経験値を取得
	var exp model.Exp
	exp,err = model.GetExp(touken.ToushuID,req.Level)
	if err != nil{
		logging.Logger.Error(fmt.Sprintf("Failed to get Exp by toukenID(%d) and level(%d)", req.Touken,req.Level))
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to get Exp by toukenIDand level in GetToukenOne"})
		return
	}

	//金平糖何個分？
	konpeto := float64(exp.SumExp) / constant.KonpetoExp

	//厚樫山何周分？
	atsukashiKari := float64(exp.SumExp) / constant.AtsukashiExp
	atsukashi := math.Round(atsukashiKari * 100) / 100

	logging.Logger.Info("Success GetToukenOne", zap.Int("toukenID",req.Touken),zap.Int("level",req.Level),zap.String("saniwa",req.Saniwa))


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