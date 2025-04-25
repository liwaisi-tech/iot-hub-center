package gorm

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	db *gorm.DB
	once sync.Once
)


func initPostgresDatabase()(db *gorm.DB, err error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	return
}


func GetPostgresDB() (postgresDB *gorm.DB, err error) {
	if db != nil{
		return db, nil
	}
	once.Do(func() {
		db, err = initPostgresDatabase()
	})
	return db, err
}