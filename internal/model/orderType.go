package model

import "gorm.io/gorm"

type OrderType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"Description"`
}
