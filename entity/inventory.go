package entity

type Inventory struct{
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Detail string `gorm:"type:text" json:"detail"`
	UserID uint64 `gorm:"not null" json:"-"`
}