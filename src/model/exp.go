package model

import "touken-exp/src/db"

//Exp expテーブルデータ
type Exp struct{
	ID string
	ToushuID int
	Level int
	SumExp int32
}

func GetExpAll()([]Exp, error){
	exp := []Exp{}
	if err := db.Conn.Find(&exp).Error; err != nil{
		return nil, err
	}
	return exp,nil

}