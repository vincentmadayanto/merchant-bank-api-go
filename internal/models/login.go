package models

// LoginRequest represents the data for the login request
type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// LoginResponse represents the data returned after login
type LoginResponse struct {
	Token string `json:"token"`
}
