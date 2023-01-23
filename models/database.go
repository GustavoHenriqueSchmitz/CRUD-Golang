package models

import (
	"time"
)

type People struct {
	ID        int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name      string `json:"name" gorm:"not null"`
	Age       int    `json:"age" gorm:"not null"`
	Function  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
