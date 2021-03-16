package handler

import (
	"log"
	"math"
	"net/http"
	"touken-exp/src/constant"
	"touken-exp/src/model"

	"github.com/gin-gonic/gin"
)

//GetToukenTwo ２振り
func GetToukenTwo(c *gin.Context){

	//リクエストボディからデータを取得
	var req getToukenTwoRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		log.Println("failed to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message":"failed!!!!!!"})
		return
	}

	devSlice := []devStruct{
		{
		ToukenID: req.Touken1,
		Level: req.Level1,
		},
		{
		ToukenID: req.Touken2,
		Level: req.Level2,
		},
	}

	var toukenDataSlice []toukenSingle 

	for _,t := range devSlice{
		//刀剣IDから刀剣名、刀種IDを取得
		touken,err := model.GetTouken(t.ToukenID)
		if err != nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message":"internal server error"})
			return
		}

		//刀種IDとレベルから経験値を取得
		var exp model.Exp
		exp,err = model.GetExp(touken.ToushuID,t.Level)
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

		toukenDataSlice = append(toukenDataSlice, toukenSingle{
			ToukenName: touken.Touken,
			Level:t.Level,
			Exp: exp.SumExp,
			Konpeto: konpeto,
			Atsukashi: atsukashi,
			},
		)
	}

	//金平糖何個分、厚樫山何周分で追いつくか,同じ刀剣だった場合の処理
	var diffKonpeto,diffAtsukashi,moreGrown,lessGrown int
	if(toukenDataSlice[0].Exp > toukenDataSlice[1].Exp){
		diffKonpeto = int((toukenDataSlice[0].Exp - toukenDataSlice[1].Exp) / constant.KonpetoExp)
		diffAtsukashi = int((toukenDataSlice[0].Exp - toukenDataSlice[1].Exp) / constant.AtsukashiExp)
		moreGrown = 0
		lessGrown = 1
		if(toukenDataSlice[0].ToukenName == toukenDataSlice[1].ToukenName){
			toukenDataSlice[0].ToukenName = "一振り目の" + toukenDataSlice[0].ToukenName
			toukenDataSlice[1].ToukenName = "二振り目の" + toukenDataSlice[1].ToukenName
		}
	}else if(toukenDataSlice[0].Exp < toukenDataSlice[1].Exp){
		diffKonpeto = int((toukenDataSlice[1].Exp - toukenDataSlice[0].Exp) / constant.KonpetoExp)
		diffAtsukashi = int((toukenDataSlice[1].Exp - toukenDataSlice[0].Exp) / constant.AtsukashiExp)
		moreGrown = 1
		lessGrown = 0
		if(toukenDataSlice[0].ToukenName == toukenDataSlice[1].ToukenName){
			toukenDataSlice[0].ToukenName += "二振り目の" + toukenDataSlice[0].ToukenName
			toukenDataSlice[1].ToukenName += "一振り目の" + toukenDataSlice[1].ToukenName
		}
	}else{
		diffKonpeto = 0
		diffAtsukashi = 0
		moreGrown = -1
		lessGrown = -1
	}

	c.JSON(http.StatusOK, getToukenTwoResponse{
		Touken: toukenDataSlice,
		MoreGrown: moreGrown,
		LessGrown: lessGrown,
		DiffKonpeto: diffKonpeto,
		DiffAtsukashi: diffAtsukashi,
	})

}

type getToukenTwoRequest struct{
	Touken1 int `json:"touken1,string,omitempty"`
	Level1 int `json:"level1,string,omitempty"`
	Touken2 int `json:"touken2,string,omitempty"`
	Level2 int `json:"level2,string,omitempty"`
	Saniwa string `json:"saniwa"`
}

type getToukenTwoResponse struct{
	Touken []toukenSingle `json:"touken"`
	MoreGrown int `json:"moreGrown"`
	LessGrown int `json:"lessGrown"`
	DiffKonpeto int `json:"diffKonpeto"`
	DiffAtsukashi int `json:"diffAtsukashi"`
}

type toukenSingle struct{
	ToukenName string `json:"toukenName"`
	Level int `json:"level"`
	Exp int32 `json:"exp"`
	Konpeto float64 `json:"konpeto"`
	Atsukashi float64 `json:"atsukashi"`	
}

type devStruct struct{
	ToukenID int
	Level int
}