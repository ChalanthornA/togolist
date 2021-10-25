package entity

import (
	"time"
)

type Inventory struct{
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Detail string `gorm:"type:text" json:"detail"`
	Price int64 `json:"price"`
	UserRefer int64
	User User `json:"user" gorm:"foreignKey:user_refer;references:id"`
	Create   time.Time `json:"create"`
	Expired time.Time 
}