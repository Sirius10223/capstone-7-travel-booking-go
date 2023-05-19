package controllers

import (
	"Capstone_Project/entity"
	"Capstone_Project/models"
	"Capstone_Project/usecases"
	"encoding/json"
	"net/http"
	"strconv"
)

type HistoryControllerInterface interface {
	ShowAllHistory(w http.ResponseWriter, r *http.Request)
	ShowAllHistoryByuserID(w http.ResponseWriter, r *http.Request)
	SaveHistory(w http.ResponseWriter, r *http.Request)
	UpdateHistory(w http.ResponseWriter, r *http.Request)
	DeleteHistory(w http.ResponseWriter, r *http.Request)
}

type HistoryController struct {
	historyUseCase usecases.HistoryUseCaseInterface
}

func NewHistoryController(historyUseCase usecases.HistoryUseCaseInterface) HistoryControllerInterface {
	return &HistoryController{historyUseCase}
}

func (hus *HistoryController) ShowAllHistory(w http.ResponseWriter, r *http.Request) { //! for admin-only
	histories, err := hus.historyUseCase.ShowAllHistory()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ResponseJson{Message: "success get all histories", Data: histories})

}

func (hus *HistoryController) ShowAllHistoryByuserID(w http.ResponseWriter, r *http.Request) {
	//! later use context to retrieve user ID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	histories, err := hus.historyUseCase.ShowAllHistoryByuserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ResponseJson{Message: "success get all histories by user ID", Data: histories})

}

func (hus *HistoryController) SaveHistory(w http.ResponseWriter, r *http.Request) {
	var newHistory models.Histories
	err := json.NewDecoder(r.Body).Decode(&newHistory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	history, err := hus.historyUseCase.SaveHistory(newHistory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ResponseJson{Message: "success save history", Data: history})
}

func (hus *HistoryController) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	var existedHistory models.Histories
	err := json.NewDecoder(r.Body).Decode(&existedHistory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}
	
	history, err := hus.historyUseCase.UpdateHistory(existedHistory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ResponseJson{Message: "success update history", Data: history})
}

func (hus *HistoryController) DeleteHistory(w http.ResponseWriter, r *http.Request) {

	historyID, err := strconv.Atoi(r.URL.Query().Get("history_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	err = hus.historyUseCase.DeleteHistory(historyID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.ResponseJson{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ResponseJson{Message: "success update history"})
}
