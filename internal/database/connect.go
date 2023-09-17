package database

import (
	"go-htmx/internal/database/models"
	"os"

	"github.com/a631807682/zerofield"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Debug bool = false
)

func Connect() {
	db_uri := os.Getenv("DATABASE_URL")

	// Connect
	conn := postgres.Open(db_uri)
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate
	dbErr := db.AutoMigrate(
		&models.Account{},
		&models.Product{},
		&models.CartItem{},
	)

	if dbErr != nil {
		panic(dbErr)
	}

	// Plugins
	db.Use(zerofield.NewPlugin())

	if Debug {
		DB = db.Debug()
	} else {
		DB = db
	}
}
