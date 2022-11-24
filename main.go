package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yildizarzu/go-gorm-restapi/db"
	"github.com/yildizarzu/go-gorm-restapi/models"
	"github.com/yildizarzu/go-gorm-restapi/routes"
)

func main() {
	fmt.Println("Hello Word")
	db.DBConnection()
	db.DB.AutoMigrate(&models.Task{})
	db.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")
	//r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Println("Starting Server t port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
