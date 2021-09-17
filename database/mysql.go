package database

import (
	"database/sql"
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

func create(name string) {

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
}
