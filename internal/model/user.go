package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"comment:用户名" json:"username"`
	DisplayName string `json:"displayName,omitempty"`
	Password    string `json:"password"`
}
