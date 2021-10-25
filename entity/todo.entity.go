package entity

import "time"

type Todo struct{
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Todo string `json:"todo"`
	Create time.Time `json:"create"`
	UserRefer int64
	User User `json:"user" gorm:"foreignKey:user_refer;references:id"`
}