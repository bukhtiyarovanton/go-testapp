package models

import (
	"time"

	//"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Todo struct {
	Model
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	UserID    uint      `json:"user_id"`
}
