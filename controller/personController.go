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
	handleError(json.NewDecoder(r.Body).Decode(&person), w)
	tx := database.Db.Create(person)
	handleError(tx.Error, w)
	if tx.Error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		handleError(json.NewEncoder(w).Encode(person), w)
	}
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Person
	database.Db.Find(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	handleError(json.NewEncoder(w).Encode(persons), w)

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
	handleError(json.NewEncoder(w).Encode(person), w)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
