package services

import (
	"errors"
	"merchant-bank-api/internal/models"
	"merchant-bank-api/internal/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticateUser(id, password string) (string, error) {
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
		return "", errors.New("Invalid Credentials")
	}

	return generateToken(customer.ID)
}

func generateToken(id string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not set in .env file")
	}

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

func ProcessPayment(customerID string, amount float64) error {
	if amount < 1 {
		return errors.New("Payment Amount Must Be More Than Zero and Cannot Be Negative Number")
	}

	transaction := models.Transaction{
		CustomerID: customerID,
		Amount:     amount,
	}
	return repositories.AddTransaction(transaction)
}
