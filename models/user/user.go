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

func NewUser(username string, email string, passw string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: passw,
	}
}

func NewUserWithID(id uint, username string, email string, passw string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: passw,
	}
}
