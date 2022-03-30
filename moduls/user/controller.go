package user

import (
	// "io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GET_USER (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	encoder := json.NewEncoder(w)

	u, _ := GetById(vars["Id"])

	encoder.Encode(u)
}


func GET (w http.ResponseWriter, r *http.Request){
	// var := mux.Vars(r)

	encoder := json.NewEncoder(w)

	getUsers, _ := GetUser()

	encoder.Encode(getUsers)
}


func Delete(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)

  encode := json.NewEncoder(w)

  deletUser := DeleteUser(vars["id"])

  encode.Encode(deletUser)
}

// func POST (w http.ResponseWriter, r *http.Request) {

// 	body, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}

// 	u := User {}

// 	json.Unmarshal(body, &u)

// 	NewUser(u)

// 	encoder.Encode()
// }