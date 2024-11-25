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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/payment", controllers.MakePayment).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
