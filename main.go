package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lhendges/simple-go-crud/controllers"

	"github.com/gorilla/mux"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandle).Methods("GET")
	router.HandleFunc("/person/", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/person/", controllers.GetAllPersons).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPersonById).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{id}", controllers.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
