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
	//call database ORM
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	//set template
	files := []string{
		"./views/masterpage.html",
		"./views/task/index.html",
	}
	//handler template parse
	htmlTemplate, err := template.ParseFiles(files...)
	//set error if html parse error
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//set map data tasks
	var tasks []models.Task
	//get table with
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

//edit method
func (controller *TaskController) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}

	files := []string{
		"./views/masterpage.html",
		"./views/task/edit_form.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var tasks models.Task
	db.Find(&tasks, params.ByName("id"))

	data := map[string]interface{}{
		"task": tasks,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

//update method
func (controller *TaskController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	id := params.ByName("id")

	var task models.Task
	db.First(&task, id)

	task.Assignee = r.FormValue("assignee")
	task.Date = r.FormValue("date")
	task.Description = r.FormValue("description")

	query := db.Save(&task)
	if query.Error != nil {
		log.Println(query.Error)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

//mark as done method
func (controller *TaskController) MarkAsDone(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	id := params.ByName("id")

	var task models.Task
	db.First(&task, id)

	task.IsDone = true

	query := db.Save(&task)
	if query.Error != nil {
		log.Println(query.Error)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *TaskController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	id := params.ByName("id")

	var task models.Task
	db.Delete(&task, id)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *TaskController) Trash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//call database ORM
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	//set template
	files := []string{
		"./views/masterpage.html",
		"./views/task/trash.html",
	}
	//handler template parse
	htmlTemplate, err := template.ParseFiles(files...)
	//set error if html parse error
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//set map data tasks
	var tasks []models.Task
	//get table with
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&tasks)

	datas := map[string]interface{}{
		"Tasks": tasks,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

func (controller *TaskController) DeleteTrash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}

	var task models.Task
	db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&task)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *TaskController) RestoreTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("db/db_task.db"), &gorm.Config{})
	if err != nil {
		panic("Terjadi kesalahan database. Error : " + err.Error())
	}
	id := params.ByName("id")
	var task models.Task
	query := db.Unscoped().Where("id = ?", id).Find(&task).Update("deleted_at", nil)
	if query.Error != nil {
		log.Println(query.Error)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
