package entity

import ()

type User struct{
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}