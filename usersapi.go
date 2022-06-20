package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"usersapi/usrdata"
	"usersapi/usrsvc"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", SaveUser).Methods("POST").Headers("Content-Type", "application/json")
	router.HandleFunc("/users", FindUsers).Methods("GET").Headers("Accept", "application/json")
	router.HandleFunc("/users/{id}", FindUser).Methods("GET").Headers("Accept", "application/json")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT").Headers("Content-Type", "application/json")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user usrdata.User
	json.Unmarshal(reqBody, &user)
	if len(strings.TrimSpace(user.Id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usrsvc.Save(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	users := usrsvc.FindAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := usrsvc.FindOne(id)
	if reflect.DeepEqual(user, usrdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := usrsvc.FindOne(id)
	if reflect.DeepEqual(user, usrdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	if user.Id != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usrsvc.Save(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := usrsvc.FindOne(id)
	if reflect.DeepEqual(user, usrdata.User{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	usrsvc.Delete(id)
}
