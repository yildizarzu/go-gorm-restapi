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
	db.DBConnection("host=postgres user=arzu password=12345678 dbname=postgres port=5432")
	db.DB.AutoMigrate(&models.Ticket{})

	db.CreateTestDB()

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/ticket_options", routes.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/ticket_options", routes.CreateTicketOptionHandler).Methods("POST")
	r.HandleFunc("/ticket/{id}", routes.GetTicketHandler).Methods("GET")
	r.HandleFunc("/ticket_options/{id}/purchases", routes.PurchaseFromTicketOptionHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./swaggerui"))
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	//fmt.Println("Starting Server t port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
