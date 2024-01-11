package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	Posts     []Post
}

type UserView struct {
	ID    uint   `json:"ID"`
	Email string `json:"Email"`
}
