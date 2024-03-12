package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	OrderedAt    time.Time
	CustomerName string `gorm:"not null;type:varchar(50)"`
	Items		 []Item
}