package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"John", "New York", "110075"},
		{"Paul", "New York", "110075"},
	}

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
