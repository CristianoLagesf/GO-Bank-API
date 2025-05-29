package app

import (
	"log"
	"net/http"

	"github.com/CristianoLagesf/GO-Bank-API/domain"
	"github.com/CristianoLagesf/GO-Bank-API/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
