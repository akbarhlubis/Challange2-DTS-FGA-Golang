package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type InDB struct {
	DB *gorm.DB
}
