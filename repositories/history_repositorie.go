package repositories

import (
	"Capstone_Project/models"

	"gorm.io/gorm"
)

type HistoryRepositoryInterface interface {
	GetAllHistory() ([]models.Histories, error)
	GetAllHistoryByuserID(userID int) ([]models.Histories, error)
	CreateHistory(newHistory models.Histories) (models.Histories, error)
	UpdateHistory(history models.Histories) (models.Histories, error)
	DeleteHistory(historyID int) error
}

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db}
}

func (hr *HistoryRepository) GetAllHistory() ([]models.Histories, error) { //! for admin-only
	var histories []models.Histories
	if err := hr.db.Find(&histories).Error; err != nil {
		return []models.Histories{}, err
	}
	return histories, nil

}

func (hr *HistoryRepository) GetAllHistoryByuserID(userID int) ([]models.Histories, error) {
	var histories []models.Histories
	if err := hr.db.Table("histories").Joins("INNER JOIN users on histories.user_id = users.id").Find(&histories, "histories.user_id = ?", userID).Error; err != nil {
		return []models.Histories{}, err
	}
	return histories, nil
}

func (hr *HistoryRepository) CreateHistory(newHistory models.Histories) (models.Histories, error) {
	if err := hr.db.Create(&newHistory).Error; err != nil {
		return models.Histories{}, err
	}

	history := newHistory
	return history, nil
}

func (hr *HistoryRepository) UpdateHistory(history models.Histories) (models.Histories, error) {
	newHistory := history

	if err := hr.db.Model(&models.Histories{}).Where("id = ? AND user_id = ?", history.ID, history.UserID).Updates(&newHistory).Error; err != nil {
		return models.Histories{}, err
	}

	return history, nil
}

func (hr *HistoryRepository) DeleteHistory(historyID int) error {

	if err := hr.db.Where("id = ?", historyID).Unscoped().Delete(&models.Histories{}).Error; err != nil {
		return err
	}

	return nil
}
