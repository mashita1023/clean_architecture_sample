package model

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" validate:"required"`
	Mail      string     `json:"mail" validate:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
