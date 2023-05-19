package models

import "gorm.io/gorm"

type Histories struct {
	gorm.Model
	Name string 
}