package models

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"not null;type:varchar(10)"`
	Description string `gorm:"not null;type:varchar(50)"`
	Quantity    uint
	OrderId     uint
}