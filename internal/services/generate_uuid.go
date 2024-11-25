package services

import "github.com/google/uuid"

// generateID generates a unique identifier for transactions
func generateID() string {
	return uuid.New().String()
}
