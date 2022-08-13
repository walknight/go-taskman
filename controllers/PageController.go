package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type PageController struct{}

func (controller *PageController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//set template
	files := []string{
		"./views/masterpage.html",
		"./views/about/index.html",
	}
	//handler template parse
	htmlTemplate, err := template.ParseFiles(files...)
	//set error if html parse error
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
