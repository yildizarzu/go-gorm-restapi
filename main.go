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
	db.DB.AutoMigrate(&models.Ticket{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/ticket_options", routes.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket_options", routes.CreateTicketOption).Methods("POST")
	r.HandleFunc("/ticket/{id}", routes.GetTicketHandler).Methods("GET")
	r.HandleFunc("/ticket_options/{id}/purchases", routes.PurchaseFromTicketOptionHandler).Methods("POST")

	fmt.Println("Starting Server t port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
