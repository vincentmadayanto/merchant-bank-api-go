package controllers

import (
	"encoding/json"
	"log"
	"merchant-bank-api/internal/models"
	"merchant-bank-api/internal/services"
	"net/http"
)

// Login handles user login requests
func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := services.AuthenticateUser(req.ID, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := models.LoginResponse{Token: token}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// Log the error and return an internal server error
		log.Println("Error writing response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// MakePayment handles payment requests
func MakePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	customerID := r.Header.Get("CustomerID")
	err := services.ProcessPayment(customerID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Payment successful")); err != nil {
		// Log the error and return an internal server error
		log.Println("Error writing response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Logout handles user logout requests
func Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Logout successful")); err != nil {
		// Log the error and return an internal server error
		log.Println("Error writing response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
