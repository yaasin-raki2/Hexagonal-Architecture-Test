package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yaasin-raki2/banking/domain"
	"github.com/yaasin-raki2/banking/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" {
		log.Fatalf("Missing environment variables")
	}
}

func Start() {
	// Ensure we have the right environment variables.
	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//Starting Server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)

	if err != nil {
		log.Fatal(err)
	}
}
