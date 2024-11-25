// services/auth_service.go
package services

import (
	"errors"
	"merchant-bank-api/internal/models"
	"merchant-bank-api/internal/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthenticateUser authenticates a user and generates a token
func AuthenticateUser(id, password string) (string, error) {
	// Simulate authentication (replace with actual DB logic)
	customers, err := repositories.LoadCustomers()
	if err != nil {
		return "", err
	}

	var customer models.Customer
	for _, c := range customers {
		if c.ID == id && c.Password == password {
			customer = c
			break
		}
	}

	if customer.ID == "" {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	return generateToken(customer.ID)
}

// generateToken generates a JWT token for the user
func generateToken(id string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not set in .env file")
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ProcessPayment processes a payment (this would normally interact with a DB or external service)
func ProcessPayment(customerID string, amount float64) error {
	// Dummy validation: if amount is below 1, fail
	if amount < 1 {
		return errors.New("payment amount is too low")
	}

	// Add the transaction to a mock DB or file
	transaction := models.Transaction{
		CustomerID: customerID,
		Amount:     amount,
	}
	return repositories.AddTransaction(transaction)
}
