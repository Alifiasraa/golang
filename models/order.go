package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(50)"`
	OrderedAt    time.Time
	Items		 []Item
}