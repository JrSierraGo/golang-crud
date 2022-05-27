package main

import (
	"crud/controller"
	"crud/database"
	"crud/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8081")

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func initHandlers(router *mux.Router) {
	router.HandleFunc("/api/person/create", controller.CreatePerson).Methods(http.MethodPost)
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
