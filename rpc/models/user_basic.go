package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
