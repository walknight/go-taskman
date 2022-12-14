package main

import (
	"go-taskman/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//port
	port := os.Getenv("PORT")
	//open database
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	//check error db
	if err != nil {
		panic("Terjadi kesalahan. Error : " + err.Error())
	}

	//migrasi data struct model ke database schema
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("Terjadi kesalahan. Error : " + err.Error())
	}

	initRouter(port)

}
