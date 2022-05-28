package controller

import (
	"crud/database"
	"crud/entity"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person entity.Person
	json.NewDecoder(r.Body).Decode(&person)
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

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Person
	database.Db.Find(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)

}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Db.Find(&person, key)
	w.Header().Set("Content-Type", "application/json")
	if person == (entity.Person{}) {
		return
	}
	json.NewEncoder(w).Encode(person)
}
