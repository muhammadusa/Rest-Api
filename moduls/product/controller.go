package product

import (
	"net/http"
	"encoding/json"
)

func GetMany (w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)

	encoder.Encode(Products)
}
