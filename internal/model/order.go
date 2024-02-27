package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User        uint   `json:"-"`
	Count       uint64 `json:"count"`
	OrderType   uint   `json:"orderType"`
	Description string `json:"description"`
}
