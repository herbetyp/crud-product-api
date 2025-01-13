package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password,omitempty" gorm:"not null"`
	UId       string         `json:"uid,omitempty" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
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
