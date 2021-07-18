package app

import (
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	// Define Routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//Starting Server
	err := http.ListenAndServe("localhost:8080", mux)
	LogIfErr(err)
}
