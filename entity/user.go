package entity

import ()

type User struct{
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
}