package handler

import (
	"fmt"
	"log"
	"net/http"
	"tkrb-exp-api/src/logging"
	"tkrb-exp-api/src/model"

	"github.com/gin-gonic/gin"
)

//GetToukenList 刀剣リストを取得する
func GetToukenList(c *gin.Context){
	logging.Logger.Info(fmt.Sprintf("Get access to %v", c.Request.URL.Path))

	toukenList,err := model.GetToukenAll()
	if err != nil{
		logging.Logger.Error("Failed to get toukenNameList")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get toukenNameList"})
		return
	}
	var toukenNameList []ToukenData
	for _,t := range toukenList{
		toukenData := ToukenData{
			Key: t.ToukenID,
			ToukenID : t.ToukenID,
			ToukenName: t.Touken,
		}
		toukenNameList = append(toukenNameList, toukenData)
	}

	logging.Logger.Info("Success GetToukenList")
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
