package main

import (
	"fmt"
	"go-taskman/controllers"
	"go-taskman/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//port
	port := ":8050"
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

	//declare controller
	taskController := &controllers.TaskController{}

	router := httprouter.New()

	router.GET("/", taskController.Index)
	router.GET("/create", taskController.Create)
	router.POST("/create", taskController.Create)

	fmt.Println("Starting web server at http://localhost" + port)
	http.ListenAndServe(port, router)

}
