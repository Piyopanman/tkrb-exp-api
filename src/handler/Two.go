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

//GetToukenTwo ２振り
func GetToukenTwo(c *gin.Context){
	logging.Logger.Info(fmt.Sprintf("Get access to %v", c.Request.URL.Path))

	//リクエストボディからデータを取得
	var req getToukenTwoRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		logging.Logger.Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"message":"Failed to bind JSON in GetToukenTwo"})
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
			logging.Logger.Error(fmt.Sprintf("Failed to get toukenName and toukenID by toukenID(%d)", t.ToukenID))
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to get toukenName and toukenID by toukenID"})
			return
		}

		//刀種IDとレベルから経験値を取得
		var exp model.Exp
		exp,err = model.GetExp(touken.ToushuID,t.Level)
		if err != nil{
			logging.Logger.Error(fmt.Sprintf("Failed to get Exp by toushuID(%d) and level(%d)", touken.ToushuID,t.Level))
			c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to get Exp by toushuID and level"})
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
	var diffKonpeto,diffAtsukashi float64
	var moreGrown,lessGrown int
	var isSameExp bool
	if(toukenDataSlice[0].Exp > toukenDataSlice[1].Exp){
		diffKonpeto = math.Ceil(float64(toukenDataSlice[0].Exp - toukenDataSlice[1].Exp) / constant.KonpetoExp)
		diffAtsukashi = math.Ceil(float64(toukenDataSlice[0].Exp - toukenDataSlice[1].Exp) / constant.AtsukashiExp)
		isSameExp = false
		moreGrown = 0
		lessGrown = 1
		if(toukenDataSlice[0].ToukenName == toukenDataSlice[1].ToukenName){
			toukenDataSlice[0].ToukenName = "一振り目の" + toukenDataSlice[0].ToukenName
			toukenDataSlice[1].ToukenName = "二振り目の" + toukenDataSlice[1].ToukenName
		}
	}else if(toukenDataSlice[0].Exp < toukenDataSlice[1].Exp){
		diffKonpeto = math.Ceil(float64(toukenDataSlice[1].Exp - toukenDataSlice[0].Exp) / constant.KonpetoExp)
		diffAtsukashi = math.Ceil(float64(toukenDataSlice[1].Exp - toukenDataSlice[0].Exp) / constant.AtsukashiExp)
		isSameExp = false
		moreGrown = 1
		lessGrown = 0
		if(toukenDataSlice[0].ToukenName == toukenDataSlice[1].ToukenName){
			toukenDataSlice[0].ToukenName = "二振り目の" + toukenDataSlice[0].ToukenName
			toukenDataSlice[1].ToukenName = "一振り目の" + toukenDataSlice[1].ToukenName
		}
	}else{
		diffKonpeto = 0
		diffAtsukashi = 0
		isSameExp = true
		moreGrown = 0
		lessGrown = 1
		if(toukenDataSlice[0].ToukenName == toukenDataSlice[1].ToukenName){
			toukenDataSlice[0].ToukenName = "一振り目の" + toukenDataSlice[0].ToukenName
			toukenDataSlice[1].ToukenName = "二振り目の" + toukenDataSlice[1].ToukenName
		}
	}

	logging.Logger.Info("Seccess GetToukenTwo", zap.Int("touken1", req.Touken1), zap.Int("touken2",req.Touken2), zap.String("saniwa",req.Saniwa))

	c.JSON(http.StatusOK, getToukenTwoResponse{
		Touken: toukenDataSlice,
		IsSameExp: isSameExp,
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
	IsSameExp bool `json:"isSameExp"`
	MoreGrown int `json:"moreGrown"`
	LessGrown int `json:"lessGrown"`
	DiffKonpeto float64 `json:"diffKonpeto"`
	DiffAtsukashi float64 `json:"diffAtsukashi"`
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