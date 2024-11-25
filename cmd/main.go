package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"merchant-bank-api/internal/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/payment", controllers.MakePayment).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
