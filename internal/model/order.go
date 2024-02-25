package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User  int64
	Count int64
}
