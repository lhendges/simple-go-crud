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
		response := MessageResponse{Message: err.Error()}
		res, _ := utils.ToJson(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
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

	if err != nil {
		res, _ := utils.ToJson(MessageResponse{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		res, _ := utils.ToJson(persons)
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

	if parseErr != nil {
		res, _ := utils.ToJson(MessageResponse{Message: parseErr.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else if err != nil {
		res, _ := utils.ToJson(MessageResponse{Message: err.Error()})
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		res, _ := utils.ToJson(foundPerson)
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
		res, _ := utils.ToJson(MessageResponse{Message: findErr.Error()})
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else if updateErr != nil {
		res, _ := utils.ToJson(MessageResponse{Message: updateErr.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
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
		res, _ := utils.ToJson(MessageResponse{Message: parseErr.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}

	foundPerson, findErr := person.GetPersonById(ID)

	if findErr != nil {
		res, _ := utils.ToJson(MessageResponse{Message: findErr.Error()})
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}

	delErr := foundPerson.DeletePerson()

	if delErr != nil {
		res, _ := utils.ToJson(MessageResponse{Message: delErr.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
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
