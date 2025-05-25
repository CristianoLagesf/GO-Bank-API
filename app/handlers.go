package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zip_code" xml:"zip_code"`
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	custormers := []Customer{
		{Name: "John", City: "Toronto", Zipcode: 15887},
		{Name: "Tim", City: "London", Zipcode: 8885877},
		{Name: "Pedro", City: "Miami", Zipcode: 72887},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(custormers)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode((custormers))
	}

}
