package database

import (
	"database/sql"
	"famesensor/go-graphql-jwt/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	godotenv.Load()
	connectDatabase()
}

func GetDB() *gorm.DB {
	return db
}

func connectDatabase() {
	configDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))

	var err error
	db, err = gorm.Open(postgres.Open(configDB), initConfig())
	if err != nil {
		panic("Connect database failed!")
	}

	err = migrate(&models.User{})
	if err != nil {
		fmt.Errorf("Migrate database error : %v", err)
		os.Exit(0)
	}
}

func initConfig() *gorm.Config {
	return &gorm.Config{}
}

func migrate(tables ...interface{}) error {
	return db.AutoMigrate(tables...)
}

func DisconnectDatabase(database *sql.DB) {
	database.Close()
}
