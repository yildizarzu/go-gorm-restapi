package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/yildizarzu/go-gorm-restapi/db"
	"github.com/yildizarzu/go-gorm-restapi/models"
	"gorm.io/gorm"
)

var mockTicket = models.Ticket{
	ID:         1,
	Name:       "Ticket Name",
	Desc:       "Ticket Description",
	Allocation: 2,
}

func TestCreateTicketOptionSuccessful(t *testing.T) {
	db.DBConnection("host=postgres user=arzu password=12345678 dbname=testpostgres port=5432")
	db.DB.Exec("DROP TABLE IF EXISTS tickets")
	db.DB.AutoMigrate(&models.Ticket{})

	body := bytes.NewBufferString(`
		{
			"name": "Ticket Name",
			"desc": "Ticket Description",
			"allocation": 2
		}
	`)

	mockRequest := httptest.NewRequest(http.MethodPost, "/ticket_options", body)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockCreateTicketOption(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	var ticket models.Ticket
	json.NewDecoder(res.Body).Decode(&ticket)

	if !reflect.DeepEqual(ticket, mockTicket) {
		t.Errorf("FAILED: expected %v, got %v\n", mockTicket, ticket)
	}
}

func TestCreateTicketOptionFailed(t *testing.T) {
	db.DBConnection("host=postgres user=arzu password=12345678 dbname=testpostgres port=5432")
	body := bytes.NewBufferString(`
		{
			"testName": "Ticket Name",
			"desc": "Ticket Description",
			"allocation": "2"
		}
	`)

	mockRequest := httptest.NewRequest(http.MethodPost, "/ticket_options", body)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockCreateTicketOption(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf(err.Error())
	}

	response := string(bodyBytes)

	if !reflect.DeepEqual(response, "Body Error") {
		t.Errorf("FAILED: expected %v, got %v\n", "Body Error", response)
	}

}

func TestGetTicketSuccessful(t *testing.T) {
	mockRequest := httptest.NewRequest(http.MethodGet, "/ticket/1", nil)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockGetTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	var ticket models.Ticket
	json.NewDecoder(res.Body).Decode(&ticket)

	if !reflect.DeepEqual(ticket, mockTicket) {
		t.Errorf("FAILED: expected %v, got %v\n", mockTicket, ticket)
	}
}

func TestGetTicketFailed(t *testing.T) {
	mockRequest := httptest.NewRequest(http.MethodGet, "/ticket/101", nil)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	vars := map[string]string{
		"id": "101",
	}

	mockRequest = mux.SetURLVars(mockRequest, vars)

	MockGetTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	var ticket models.Ticket
	json.NewDecoder(res.Body).Decode(&ticket)

	if !reflect.DeepEqual(ticket, mockTicket) {
		t.Errorf("FAILED: expected %v, got %v\n", mockTicket, ticket)
	}
}

func TestTicketAllocation(t *testing.T) {
	mockRequest := httptest.NewRequest(http.MethodGet, "/ticket/1", nil)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	mockRequest = mux.SetURLVars(mockRequest, vars)

	MockGetTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	var ticket models.Ticket
	json.NewDecoder(res.Body).Decode(&ticket)

	if !reflect.DeepEqual(ticket.Allocation, 0) {
		t.Errorf("FAILED: expected %v, got %v\n", 0, ticket.Allocation)
	}
}
func TestTicketPurchaseSuccessful(t *testing.T) {
	body := bytes.NewBufferString(`
		{
			"quantity": 2,
			"user_id": "123456"
		}
	`)

	mockRequest := httptest.NewRequest(http.MethodPost, "/ticket_options/1/purchase", body)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockPurchaseTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf(err.Error())
	}

	response := string(bodyBytes)

	if !reflect.DeepEqual(response, "Purchase Complete") {
		t.Errorf("FAILED: expected %v, got %v\n", "Purchase Complete", response)
	}
}

func TestTicketPurchaseFailed(t *testing.T) {
	body := bytes.NewBufferString(`
		{
			"quantity": 2,
			"user_id": "123456"
		}
	`)

	mockRequest := httptest.NewRequest(http.MethodPost, "/ticket_options/1/purchase", body)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockPurchaseTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf(err.Error())
	}

	response := string(bodyBytes)

	if !reflect.DeepEqual(response, "Not available ticket allocation") {
		t.Errorf("FAILED: expected %v, got %v\n", "Not available ticket allocation", response)
	}
}

func TestGetTickets(t *testing.T) {
	mockRequest := httptest.NewRequest(http.MethodGet, "/ticket_options", nil)
	mockRequest.Header.Set("Content-Type", "application/json")
	mockWriter := httptest.NewRecorder()

	MockGetTicket(mockWriter, mockRequest)

	res := mockWriter.Result()
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf(err.Error())
	}

	response := string(bodyBytes)

	if reflect.DeepEqual(response, "Ticket not found") {
		t.Errorf("FAILED: expected %v, got %v\n", "Ticket not found", response)
	}
}

func MockCreateTicketOption(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)

	if ticket.Name == "" || ticket.Desc == "" || ticket.Allocation == 0 {
		w.Write([]byte("Body Error"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println("Test Success" + ticket.Name)
		createdTicket := db.DB.Create(&ticket)
		err := createdTicket.Error

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(&ticket)
		w.WriteHeader(http.StatusOK)
	}
}

func MockPurchaseTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket models.Ticket
	result := db.DB.First(&ticket, params["id"])

	if ticket.ID == 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.Write([]byte("Ticket not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
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
}

func MockGetTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket models.Ticket
	db_Result := db.DB.First(&ticket, params["id"]).Find(&ticket)

	if ticket.ID == 0 || errors.Is(db_Result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket not found"))
		return
	} else {
		json.NewEncoder(w).Encode(&ticket)
		w.WriteHeader(http.StatusOK)
	}
}

func MockGetTickets(w http.ResponseWriter, r *http.Request) {
	var tickets []models.Ticket
	result := db.DB.Find(&tickets)
	if result.RowsAffected == 0 {
		w.Write([]byte("Ticket not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(&tickets)
		w.WriteHeader(http.StatusOK)
	}

}
