package main

import (
	"Capstone_Project/configs"
	"os"
)

func main() {

	os.Setenv("DATABASE_URL", "postgres://postgres:12345@localhost:5432/postgres")
	err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	// db := configs.GetDBConnection()


	
}