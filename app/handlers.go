package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/sancroth/mux-api/service"
	"net/http"
)

type Customer struct{
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		_ = xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(customers)
	}

	fmt.Print(w,customers)
}
