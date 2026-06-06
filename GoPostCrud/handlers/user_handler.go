package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asim9115/GoLang/models"
	"github.com/asim9115/GoLang/repository"

	//"github.com/asim9115/GoLang/utils"

	"github.com/gorilla/mux"
)

//POST /users

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
    if err := repository.CreateUser(&user); err != nil {
		
			log.Println(err)
    http.Error(w, "Failed to create user", http.StatusInternalServerError)
    return
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)

	
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users,err := repository.GetAllUsers()

	if err != nil {
		
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		
		return
	}
	json.NewEncoder(w).Encode(users)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	user, err := repository.GetUserById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	id := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = repository.UpdateUser(&user, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if err := repository.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	return
	}
	w.WriteHeader(http.StatusNoContent)


}