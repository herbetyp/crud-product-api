package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null"`
	Password  string         `json:"password" gorm:"not null"`
	UId       string         `json:"uid" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastLogin time.Time      `json:"last_login"`
}
