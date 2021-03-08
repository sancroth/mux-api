package app

import (
	"github.com/gorilla/mux"
	"github.com/sancroth/mux-api/domain"
	"github.com/sancroth/mux-api/service"
	"log"
	"net/http"
)

func Start(){
	mux := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000",mux))
}