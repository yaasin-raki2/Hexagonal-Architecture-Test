package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaasin-raki2/banking/domain"
	"github.com/yaasin-raki2/banking/service"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//Starting Server
	err := http.ListenAndServe("localhost:8080", router)

	if err != nil {
		log.Fatal(err)
	}
}
