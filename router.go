package main

import (
	"fmt"
	"go-taskman/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initRouter(port string) {
	//declare controller
	taskController := &controllers.TaskController{}
	pageController := &controllers.PageController{}

	router := httprouter.New()

	//set untuk bisa baca static folder css, js, image dari folder public
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.GET("/", taskController.Index)
	router.GET("/create", taskController.Create)
	router.POST("/create", taskController.Create)
	router.GET("/edit/:id", taskController.Edit)
	router.POST("/update/:id", taskController.Update)
	router.POST("/mark-as-done/:id", taskController.MarkAsDone)
	router.POST("/delete/:id", taskController.Delete)
	router.GET("/about", pageController.Index)
	router.GET("/trash", taskController.Trash)
	router.POST("/delete-trash", taskController.DeleteTrash)
	router.POST("/restore-task/:id", taskController.RestoreTask)

	if port == "" {
		port = "8050" // Default port if not specified
	}

	fmt.Println("Starting web server at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
