package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TaskController struct{}

func (controller *TaskController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}

	files := []string{
		"./views/masterpage.html",
		"./views/task/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datas := map[string]interface{}{}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
