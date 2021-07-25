package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/yaasin-raki2/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		contentType = "application/json"
	}

	w.Header().Add("Content-Type", contentType)

	if contentType == "application/xml" {
		xml.NewEncoder(w).Encode(customers)
	} else if contentType == "application/json" {
		json.NewEncoder(w).Encode(customers)
	}
}
