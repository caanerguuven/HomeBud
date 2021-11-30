package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "template1"
)

var database *gorm.DB

//Connect provides connection
func Connect() {
	var _database = GetDB()
	if _database != nil {
		fmt.Println("DB Connection is already open")
		return
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	_database, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = _database.DB().Ping()
	if err != nil {
		panic(err)
	}

	database = _database

	fmt.Println("DB Connection is successful")
	//defer _database.Close()
}

//GetDB gets the current db
func GetDB() *gorm.DB {
	return database
}
