package models

import "gorm.io/gorm"

type Histories struct {
	gorm.Model
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
