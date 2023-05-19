package repositories

import "github.com/jinzhu/gorm"

type HistoryRepositoryInterface interface{}

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository (db *gorm.DB) *HistoryRepository{
	return &HistoryRepository{db}
}

 