package main

import (
	"log"
	"github.com/asim9115/GoLang/config"
	"github.com/asim9115/GoLang/handlers"
	"github.com/asim9115/GoLang/models"



	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	
	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}