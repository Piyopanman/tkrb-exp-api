package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Hito 一振り用
func Hito(c *gin.Context){
	fmt.Println("Hito is called")
}