package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB // база данных

func init() {
	e := godotenv.Load() // загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("host")
	
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",dbHost,userName,dbName,password) // создать строку подключения
	fmt.Println(dbUri)

	conn,err := gorm.Open("postgres",dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{},&Contact{}) // Миграция баз данных
}

// возвращает дескприптор объекта DB
func GetDB() *gorm.DB {
	return db
}

