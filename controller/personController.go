package controller

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	tx := database.Db.Create(person)
	w.Header().Set("Content-Type", "application/json")
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(tx.Error.Error())
		json.NewEncoder(w).Encode(tx.Error.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(person)
	}
}
