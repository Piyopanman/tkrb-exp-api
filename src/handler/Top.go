package handler

import (
	"log"
	"touken-exp/src/model"

	"github.com/gin-gonic/gin"
)

//Top トップ画面
func Top(c *gin.Context){
	log.Println("Top.go is called!")

	toukens,err := model.GetToukenAll()
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(toukens)

	toushu,err := model.GetToushuAll()
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(toushu)

	exp,err := model.GetExpAll()
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(exp)


}