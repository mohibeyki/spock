package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mohibeyki/spock/model"
	"github.com/mohibeyki/spock/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db    *gorm.DB
	err   error
	dbErr error
)

// Init opens a database and saves the reference to `Database` struct.
func Init() {
	config := config.GetConfig()

	host := config.Database.Host
	port := config.Database.Port
	dbName := config.Database.Dbname
	username := config.Database.Username
	password := config.Database.Password

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, username, password, dbName, port, "disable")
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		log.Fatalln("db err: ", err)
	}

	db.AutoMigrate(&model.User{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}

// GetDBErr returns the dbErr instance
func GetDBErr() error {
	return dbErr
}
