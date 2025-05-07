package api

import (
	"encoding/json"
	"fmt"
	"leagueinform/internal/types"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type APIHandler struct {
	db *gorm.DB
}

func (h *APIHandler) createAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var acc types.Account

		err := json.NewDecoder(req.Body).Decode(&acc)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		rs := h.db.Where(&types.Account{Name: acc.Name, Tag: acc.Tag}).Attrs(&types.Account{Name: acc.Name, Tag: acc.Tag}).FirstOrCreate(&acc)
		if rs.Error != nil {
			log.Println("Query error:", rs.Error)
		}

		// rs := h.db.Create(&acc)

	}

}

func (h *APIHandler) login(w http.ResponseWriter, req *http.Request) {

}

func RunServer(DB *gorm.DB) {
	handler := &APIHandler{db: DB}

	http.HandleFunc("/hey2", handler.createAccount)
	fmt.Println("test")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
