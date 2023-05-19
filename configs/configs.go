package configs

import (
	"Capstone_Project/models"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB


type Config struct {
	
}

func InitDB() error {
	// connect using gorm pgx
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:        os.Getenv("DATABASE_URL"),
		DriverName: "pgx",
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(models.User{},models.Histories{})
	SetupDBConnection(conn)
	return nil
}

// store in a new var
func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

// get db obj
func GetDBConnection() *gorm.DB {
	return db
}

