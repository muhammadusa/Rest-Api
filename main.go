package main

import (
	"net/http"
	"app/moduls/user"
	"app/moduls/product"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/products", product.GetMany).Methods("GET")
	r.HandleFunc("/users/{Id}", user.GET_USER).Methods("GET")
	r.HandleFunc("/users", user.GET).Mehtods("GET")
	r.HandleFunc("/users/{Id}", user.DELETE).Methods("DELETE")
	// r.HandleFunc("/users", user.POST)
	// r.HandleFunc("/users", )


	http.ListenAndServe(":8080", r)
}