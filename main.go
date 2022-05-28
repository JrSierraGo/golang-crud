package main

import (
	"crud/controller"
	"crud/database"
	"crud/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8081")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func initHandlers(router *mux.Router) {
	personRouter := router.PathPrefix("/api/person").Subrouter()
	personRouter.HandleFunc("/create", controller.CreatePerson).Methods(http.MethodPost)
	personRouter.HandleFunc("/{id:[0-9]+}", controller.GetPersonById).Methods(http.MethodGet)
	personRouter.HandleFunc("/list", controller.GetAllPeople).Methods(http.MethodGet)
}

func initDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "root",
		DB:         "crud"}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Person{})
}
