package model

import "tkrb-exp-api/src/db"

//Toushu 刀種テーブルデータ
type Toushu struct{
	ToushuID int
	Toushu string
}

//GetToushuAll すべて
func GetToushuAll()([]Toushu, error){
	toushu := []Toushu{}
	if err := db.Conn.Find(&toushu).Error; err != nil{
		return nil, err
	}
	return toushu,nil
}