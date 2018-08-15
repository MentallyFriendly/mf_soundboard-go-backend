package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// Init func
func Init(seed, migrate bool) *gorm.DB {
	for i := 0; i < 10; i++ {
		db, err = gorm.Open("postgres", os.Getenv("DB_PARAMS"))
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		fmt.Println("Error connecting to DB", err)
	}

	if err := db.DB().Ping(); err != nil {
		fmt.Println("Error pinging DB", err)
	}

	fmt.Println("Successfully connected to db")

	if os.Getenv("SEED") == "true" {
		seedDatabase()
		fmt.Println("Seeding database...")
	}

	if os.Getenv("MIGRATE") == "true" {
		migrateDatabase()
		fmt.Println("Migrating database...")
	}

	return db
}
