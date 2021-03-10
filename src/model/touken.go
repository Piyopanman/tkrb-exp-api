package model

import "touken-exp/src/db"

//Touken toukenテーブルデータ
type Touken struct{
	ToukenID int 
	ToushuID int 
	Touken string 
}

//GetToukenAll すべての刀剣情報を取得
func GetToukenAll()([]Touken, error){
	toukens := []Touken{}
	if err := db.Conn.Find(&toukens).Error; err != nil{
		return nil, err
	}
	return toukens,nil
}