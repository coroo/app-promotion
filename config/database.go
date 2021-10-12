package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	_ "github.com/joho/godotenv/autoload"
)

type Database struct {
	Conn *gorm.DB
	ConnSy *gorm.DB
	Err error
	ErrSy error
}

var database = Database{}

func ConnectDB() (c *gorm.DB, err error) {
	DB_CONNECTION := os.Getenv("DB_CONNECTION")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	DB_TEST := os.Getenv("DB_TEST")
	DB_DETAIL := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?parseTime=true"
	if DB_CONNECTION == "" {
		DB_DETAIL = DB_TEST
		conn, err := gorm.Open(sqlite.Open("./../../test.db"), &gorm.Config{})
		if err != nil || conn == nil {
			fmt.Println("Error connecting to DB")
			fmt.Println(err.Error())
		}
		return conn, err
	} else {
		if  database.Conn == nil{
			database.Conn, database.Err = gorm.Open(mysql.Open(DB_DETAIL), &gorm.Config{})
			if database.Err != nil || database.Conn == nil {
				fmt.Println("Error connecting to DB")
				fmt.Println(database.Err.Error())
			}
		}
		return database.Conn, database.Err
	}
}
