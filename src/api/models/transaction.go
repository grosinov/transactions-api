package models

import (
	"time"
)

type Transaction struct {
	Id       uint64    `json:"id" gorm:"primaryKey"`
	UserId   uint64    `json:"user_id"`
	Amount   float64   `json:"amount"`
	Datetime time.Time `json:"datetime"`
}
