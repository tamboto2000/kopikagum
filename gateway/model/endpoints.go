package model

import "gorm.io/gorm"

// Endpoint contains endpoint destination
type Endpoint struct {
	Host         string      `json:"host" gorm:"varchar(50);not null"`
	Port         string      `json:"port" gorm:"varchar(10)"`
	Method       string      `json:"method" gorm:"varchar(10);not null"`
	Path         string      `json:"path" gorm:"varchar(100)"`
	MiddlewareID uint        `json:"middlewareId"`
	Middleware   *Middleware `json:"middleware"`

	*gorm.Model
}
