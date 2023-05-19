package usecases

import (
	"Capstone_Project/models"
	"Capstone_Project/repositories"
)

type HistoryUseCaseInterface interface{
	ShowAllHistory  () ([]models.Histories,error)
	ShowAllHistoryByuserID  (userID int) ([]models.Histories,error) 
	SaveHistory (newHistory models.Histories) (models.Histories,error) 
	UpdateHistory (history models.Histories) (models.Histories,error)
	DeleteHistory (historyID int) error 
}

type HistoryUseCase struct {
	historyRepo repositories.HistoryRepositoryInterface
}

func NewHistoryUseCase (historyRepo repositories.HistoryRepositoryInterface) HistoryUseCaseInterface {
	return &HistoryUseCase{historyRepo}
}


func (hus *HistoryUseCase) ShowAllHistory  () ([]models.Histories,error) { //! for admin-only 
	histories,err := hus.historyRepo.GetAllHistory()
	if err != nil {
		return []models.Histories{},err
	}
	return histories, nil

}

func (hus *HistoryUseCase) ShowAllHistoryByuserID  (userID int) ([]models.Histories,error) {
	histories, err := hus.historyRepo.GetAllHistoryByuserID(userID)
	if err != nil {
		return []models.Histories{},err
	}

	return histories, nil
}

func (hus *HistoryUseCase) SaveHistory (newHistory models.Histories) (models.Histories,error) {
	history,err:= hus.historyRepo.CreateHistory(newHistory)
	if err != nil {
		return models.Histories{},err
	}
		return history, nil
}

func (hus *HistoryUseCase) UpdateHistory (history models.Histories) (models.Histories,error) {

	 history,err := hus.historyRepo.UpdateHistory(history)
	 if err != nil {
		return models.Histories{},err
	 }

	return history, nil
} 

func (hus *HistoryUseCase) DeleteHistory (historyID int) error {

	err:= hus.historyRepo.DeleteHistory(historyID)
	if err != nil {
		return err
	}
	return nil
} 



 