package routes

import (
	"encoding/json"
	"net/http"

	"github.com/yildizarzu/go-gorm-restapi/db"
	"github.com/yildizarzu/go-gorm-restapi/models"
)

func GetDenemeHandler(w http.ResponseWriter, r *http.Request) {
	var denemes []models.Deneme
	db.DB.Find(&denemes)
	json.NewEncoder(w).Encode(&denemes)
}
