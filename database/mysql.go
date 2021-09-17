package database

import (
	"fmt"
	"log"

	"github.com/AchmadAlli/go_auth_api/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	username := helper.GetEnv("DB_USERNAME")
	password := helper.GetEnv("DB_PASSWORD")
	host := helper.GetEnv("DB_HOST")
	port := helper.GetEnv("DB_PORT")
	database := helper.GetEnv("DB_DATABASE")

	destination := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true", username, password, host, port, database)

	db, err := gorm.Open("mysql", destination)
	if err != nil {
		return nil, err
	}

	log.Println("Database connected!")

	db = db.Set("gorm:table_options", "DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB")
	db = db.Set("gorm:auto_preload", true)

	return db, nil
}