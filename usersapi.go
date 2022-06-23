package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shpboris/usersdata"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"usersapi/userssvc"
)

const (
	contentTypeHeader = "Content-Type"
	acceptHeader      = "Accept"
	applicationJson   = "application/json"
	POST              = "POST"
	GET               = "GET"
	PUT               = "PUT"
	DELETE            = "DELETE"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", SaveUser).Methods(POST).Headers(contentTypeHeader, applicationJson)
	router.HandleFunc("/users", FindUsers).Methods(GET).Headers(acceptHeader, applicationJson)
	router.HandleFunc("/users/{id}", FindUser).Methods(GET).Headers(acceptHeader, applicationJson)
	router.HandleFunc("/users/{id}", UpdateUser).Methods(PUT).Headers(contentTypeHeader, applicationJson)
	router.HandleFunc("/users/{id}", DeleteUser).Methods(DELETE)
	fmt.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user usersdata.User
	json.Unmarshal(reqBody, &user)
	if len(strings.TrimSpace(user.Id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userssvc.Save(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	users := userssvc.FindAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := userssvc.FindOne(id)
	if reflect.DeepEqual(user, usersdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := userssvc.FindOne(id)
	if reflect.DeepEqual(user, usersdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	if user.Id != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userssvc.Save(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := userssvc.FindOne(id)
	if reflect.DeepEqual(user, usersdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	userssvc.Delete(id)
}
