package model

import "gorm.io/gorm"

// Middleware contains middleware destination
type Middleware struct {
	Host string `json:"host" gorm:"varchar(50);not null"`
	Port string `json:"port" gorm:"varchar(10)"`
	Path string `json:"path" gorm:"varchar(100)"`

	*gorm.Model
}
