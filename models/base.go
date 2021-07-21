package models

import "time"

type Model struct {
	ID        string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
