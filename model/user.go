package model

import (
	"BE/handlers"

	"gorm.io/gorm"
)

type Users struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}

func (Users) TableName() string {
	return "Users"
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := handlers.HashingPass(user.Password)
	user.Password = hashedPass
	return
}