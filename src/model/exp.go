package model

import "touken-exp/src/db"

//Exp expテーブルデータ
type Exp struct{
	ID int
	ToushuID int
	Level int
	SumExp int32
}


//GetExpAll すべての経験値を取得
func GetExpAll()([]Exp, error){
	exp := []Exp{}
	if err := db.Conn.Find(&exp).Error; err != nil{
		return nil, err
	}
	return exp,nil
}

//GexExp 刀種IDとレベルから経験値を取得
func GetExp(toushuID int, level int)(Exp, error){
	exp := Exp{}
	dummy := Exp{
		ID: 0,
		ToushuID: 0,
		Level: 0,
		SumExp: 0,
	}
	if err := db.Conn.Where("toushu_id=? and level=?", toushuID, level).First(&exp).Error; err!= nil{
		return dummy,err
	}
	return exp,nil
}