package app

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/home",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		})
	http.HandleFunc("/customers", getAllCustomer)

	http.ListenAndServe("localhost:8000", nil)
}
