package model

import "tkrb-exp-api/src/db"

//Touken toukenテーブルデータ
type Touken struct{
	ToukenID int 
	ToushuID int 
	Touken string 
}

//GetToukenAll すべての刀剣情報を取得
func GetToukenAll()([]Touken, error){
	toukens := []Touken{}
	if err := db.Conn.Order("touken_id").Find(&toukens).Error; err != nil{
		return nil, err
	}
	return toukens,nil
}

//GetTouken touken_idを元に情報を取得
func GetTouken(id int)(Touken, error){
	touken := Touken{}
	dummy := Touken{
		ToukenID: 0,
		ToushuID: 0,
		Touken: "",
	}
	if err := db.Conn.Where("touken_id=?", id).First(&touken).Error; err != nil{
		return dummy,err
	}
	return touken,nil
}