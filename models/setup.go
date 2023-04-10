package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3","test.db")
	
	//Enable viper to read Enviroment virables
	viper.AutomaticEnv()

	// To get the value from the config file using key
	
	//viper packege read .env

	viper_user := viper.Get("postgres")
	viper_password := viper.Get("postgres")
	viper_db := viper.Get("PokerGame")
	viper_host := viper.Get("localhost")
	viper_port := viper.Get("5434")

	postgres_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v passowrd=%v sslmode=disabled",viper_host,viper_port,viper_user,viper_db,viper_password)

	fmt.Println("conname is\t\t", postgres_conname)

	db, err := gorm.Open("postgres",postgres_conname)
	if err != nil {
		panic("Failed to connect to database!:")
	}
	db.AutoMigrate(&User{})

	// Initialise value
	m := User{Name: "Name",Password: "password",Email: "email@email.email",Phone: "812345678"}
	
	db.Create(&m)

	return db
}   