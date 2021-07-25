package app

import (
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

	//Starting Server
	err := http.ListenAndServe("localhost:8080", router)
	LogIfErr(err)
}
