package orm

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, dbport, user, password, dbname)
	gormdb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormdb, nil
}
