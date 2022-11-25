package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yildizarzu/go-gorm-restapi/db"
	"github.com/yildizarzu/go-gorm-restapi/models"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var tickets []models.Ticket
	db.DB.Find(&tickets)
	json.NewEncoder(w).Encode(&tickets)
}

func GetTicketHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket models.Ticket
	db.DB.First(&ticket, params["id"])

	if ticket.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket not found"))
		return
	}

	json.NewEncoder(w).Encode(&ticket)
}

func CreateTicketOption(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)
	createdTicket := db.DB.Create(&ticket)
	err := createdTicket.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ticket)
	w.WriteHeader(http.StatusOK)
}

func PurchaseFromTicketOptionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket models.Ticket
	db.DB.First(&ticket, params["id"])

	var ticketPurchase models.Ticket_Purchase
	json.NewDecoder(r.Body).Decode(&ticketPurchase)

	if ticket.Allocation >= ticketPurchase.Quantity {

		db.DB.Model(&models.Ticket{}).Where("id = ?", ticket.ID).Update("allocation", ticket.Allocation-ticketPurchase.Quantity)
		w.Write([]byte("Purchase Complete"))
		w.WriteHeader(http.StatusOK)
	} else {
		w.Write([]byte("Not available ticket allocation"))
		w.WriteHeader(http.StatusNotFound)
	}

}
