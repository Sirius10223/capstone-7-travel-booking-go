package main

import (
	"Capstone_Project/configs"
	"Capstone_Project/controllers"
	"Capstone_Project/repositories"
	"Capstone_Project/routes"
	"Capstone_Project/usecases"
	"net/http"
	"os"
	"sync"

	"gorm.io/gorm"
)

func main() {

	os.Setenv("DATABASE_URL", "postgres://edwin:edwinSuek@localhost:5432/travelapp")
	err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	db := configs.GetDBConnection()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		mux := http.NewServeMux()
		mux = RunServer(db, mux)

		//listener
		err = http.ListenAndServe(":8000", mux)
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()

}

func RunServer(db *gorm.DB, mux *http.ServeMux) *http.ServeMux {
	//!history service
	historyRepo := repositories.NewHistoryRepository(db)
	historyUseCase := usecases.NewHistoryUseCase(historyRepo)
	historyController := controllers.NewHistoryController(historyUseCase)

	//history handlers/controllers
	routes.MuxRoute(mux, "GET", "/api/v1/showAllHistories", http.HandlerFunc(historyController.ShowAllHistory))
	routes.MuxRoute(mux, "GET", "/api/v1/showHistory/on", http.HandlerFunc(historyController.ShowAllHistoryByuserID), "?user_id=")
	routes.MuxRoute(mux, "POST", "/api/v1/saveHistory", http.HandlerFunc(historyController.SaveHistory))
	routes.MuxRoute(mux, "PUT", "/api/v1/updateHistory", http.HandlerFunc(historyController.UpdateHistory))
	routes.MuxRoute(mux, "DELETE", "/api/v1/deleteHistory", http.HandlerFunc(historyController.DeleteHistory), "?history_id=")

	return mux

}
