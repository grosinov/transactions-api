package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	ID       uint64    `json:"id" gorm:"primaryKey"`
	UserId   uint64    `json:"user_id"`
	Amount   float64   `json:"age"`
	Datetime time.Time `json:"datetime"`
}
