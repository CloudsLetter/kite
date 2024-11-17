package db

import (
	"gorm.io/gorm"
)

type Images struct {
	ID       uint   `gorm:"primary_key"`
	Hash     string `gorm:"type:varchar(255);unique_index"`
	FileName string
	Byte     string
	gorm.Model
}
