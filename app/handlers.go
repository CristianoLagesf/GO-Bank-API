package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/CristianoLagesf/GO-Bank-API/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zip_code" xml:"zip_code"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// custormers := []Customer{
	// 	{Name: "John", City: "Toronto", Zipcode: 15887},
	// 	{Name: "Tim", City: "London", Zipcode: 8885877},
	// 	{Name: "Pedro", City: "Miami", Zipcode: 72887},
	// }
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode((customers))
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "teste post")
}
