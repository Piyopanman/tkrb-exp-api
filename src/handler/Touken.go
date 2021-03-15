package handler

import (
	"log"
	"net/http"
	"touken-exp/src/model"

	"github.com/gin-gonic/gin"
)

//刀剣リストを取得する
func GetToukenList(c *gin.Context){
	toukenList,err := model.GetToukenAll()
	if err != nil{
		log.Println("failed to get toukenNameList in GetToukenList")
		return
	}
	// log.Println(toukenList)
	var toukenNameList []ToukenData
	for _,t := range toukenList{
		toukenData := ToukenData{
			Key: t.ToukenID,
			ToukenID : t.ToukenID,
			ToukenName: t.Touken,
		}
		toukenNameList = append(toukenNameList, toukenData)
	}
	c.JSON(http.StatusOK, getToukenListResponse{List: toukenNameList})

}

type ToukenData struct{
	Key int `json:"key"`
	ToukenID int `json:"toukenID"`
	ToukenName string `json:"toukenName"`
}

type getToukenListResponse struct{
	List []ToukenData `json:"list"`
}

func GetResult(c *gin.Context){
	toukenName := c.Query("toukenName")
	if toukenName == "" {
		log.Println("toukenName is empty")
		return
	}
	level := c.Query("level")
	if level == "" {
		log.Println("level is empty")
		return
	}

}


type getResultResponse struct{
	ToukenName string `json:"toukenName"`
	Level int `json:"level"`
	Exp int32 `json:"exp"`
	Konpeto int32 `json:"konpeto"`
	Atsukashi int32 `json:"atsukashi"`
}