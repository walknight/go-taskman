package controllers

import (
	"go-taskman/models"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TaskController struct{}

// Index method
func (controller *TaskController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
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
	var tasks []models.Task
	db.Find(&tasks)

	datas := map[string]interface{}{
		"Tasks": tasks,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

//create method
func (controller *TaskController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}

	if r.Method == "POST" {
		task := models.Task{
			Assignee:    r.FormValue("assignee"),
			Date:        r.FormValue("date"),
			Description: r.FormValue("description"),
		}

		query := db.Create(&task)
		if query.Error != nil {
			log.Println(query.Error)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		files := []string{
			"./views/masterpage.html",
			"./views/task/add_form.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var tasks []models.Task
		db.Find(&tasks)

		datas := map[string]interface{}{
			"tasks": tasks,
		}

		err = htmlTemplate.ExecuteTemplate(w, "base", datas)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

}
