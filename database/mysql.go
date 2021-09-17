package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AchmadAlli/go_auth_api/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	username = helper.GetEnv("DB_USERNAME")
	password = helper.GetEnv("DB_PASSWORD")
	host     = helper.GetEnv("DB_HOST")
	port     = helper.GetEnv("DB_PORT")
)

func ConnectMysql() (*gorm.DB, error) {
	database := helper.GetEnv("DB_DATABASE")

	destination := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true", username, password, host, port, database)

	db, err := gorm.Open("mysql", destination)
	if err != nil {
		return nil, err
	}

	db = db.Set("gorm:table_options", "DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB")
	db = db.Set("gorm:auto_preload", true)

	return db, nil
}

func CreateDatabase(name string) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/", username, password, host, port)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}

	db.Close()
}

func HandleMigration(db *gorm.DB, model interface{}) {
	err := db.DropTableIfExists(model).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.AutoMigrate(model).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
