package model

import "gorm.io/gorm"

type OrderType struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OrderKindId uint      `json:"orderKindId"`
	OrderKind   OrderKind `gorm:"foreignKey:OrderKindId"`
}

type OrderKind struct {
	gorm.Model
	Name      string `json:"name"`
	Necessary bool   `json:"necessary"`
}
