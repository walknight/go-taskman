package main

import (
	"fmt"
	"go-taskman/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initRouter(port string) {
	//declare controller
	taskController := &controllers.TaskController{}

	router := httprouter.New()

	router.GET("/", taskController.Index)
	router.GET("/create", taskController.Create)
	router.POST("/create", taskController.Create)
	router.GET("/edit/:id", taskController.Edit)
	router.POST("/update/:id", taskController.Update)
	router.POST("/mark-as-done/:id", taskController.MarkAsDone)
	router.POST("/delete/:id", taskController.Delete)

	if port == "" {
		port = "8050" // Default port if not specified
	}

	fmt.Println("Starting web server at http://localhost:" + port)
	http.ListenAndServe(":"+port, router)
}
