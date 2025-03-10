package mysql_driver

import (
	"fmt"
	"log"

	"github.com/amdrx480/go-lms/drivers/mysql/categories"
	"github.com/amdrx480/go-lms/drivers/mysql/chapters"
	"github.com/amdrx480/go-lms/drivers/mysql/courses"
	"github.com/amdrx480/go-lms/drivers/mysql/modules"
	"github.com/amdrx480/go-lms/drivers/mysql/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *DBConfig) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&users.User{}, &courses.Course{}, &modules.Module{}, &chapters.Chapter{}, &categories.Category{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
