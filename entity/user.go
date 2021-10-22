package entity

import (
	"github.com/satori/go.uuid"
)

type User struct{
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}