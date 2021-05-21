package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lhendges/simple-go-crud/utils"

	person "github.com/lhendges/simple-go-crud/model"

	"github.com/gorilla/mux"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	newPerson := &person.Person{}
	utils.FromJson(r, newPerson)

	id, err := newPerson.CreatePerson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
	} else {
		response := MessageResponse{Message: fmt.Sprintf("Registro %d inserido com sucesso", id)}
		res, _ := utils.ToJson(response)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	var persons []person.Person
	persons, err := person.GetAllPersons()

	res, _ := utils.ToJson(persons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedId := vars["id"]
	ID, parseErr := strconv.ParseInt(requestedId, 0, 0)

	foundPerson, err := person.GetPersonById(ID)

	res, _ := utils.ToJson(foundPerson)
	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedId := vars["id"]
	ID, _ := strconv.ParseInt(requestedId, 0, 0)

	foundPerson, findErr := person.GetPersonById(ID)

	newPerson := &person.Person{}
	utils.FromJson(r, newPerson)

	updatedPerson, updateErr := foundPerson.UpdatePerson(newPerson.Name, newPerson.Age)

	res, _ := utils.ToJson(updatedPerson)

	if findErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(findErr.Error()))
	} else if updateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(updateErr.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedId := vars["id"]
	ID, parseErr := strconv.ParseInt(requestedId, 0, 0)

	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(parseErr.Error()))
		return
	}

	foundPerson, findErr := person.GetPersonById(ID)

	if findErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(findErr.Error()))
		return
	}

	delErr := foundPerson.DeletePerson()

	if delErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(delErr.Error()))
		return
	}

	res, _ := utils.ToJson(foundPerson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

type MessageResponse struct {
	Message string
}
