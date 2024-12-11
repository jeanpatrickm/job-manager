package config

import (
	"os"

	"github.com/jeanpatrickm/job-manager/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database does not exist, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	
	// Create and connect to the database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}
	
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}
	
	// return database
	return db, nil
}
