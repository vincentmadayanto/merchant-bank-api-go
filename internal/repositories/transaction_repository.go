// repositories/transaction_repository.go
package repositories

import (
	"encoding/json"
	"io/ioutil"
	"merchant-bank-api/internal/models"
	"os"
)

var (
	customerDataPath    = "internal/repositories/data/customers.json"
	transactionDataPath = "internal/repositories/data/transactions.json"
)

// LoadCustomers loads customers from the file (if necessary for your app)
func LoadCustomers() ([]models.Customer, error) {
	file, err := os.Open(customerDataPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var customers []models.Customer
	err = json.NewDecoder(file).Decode(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// AddTransaction adds a transaction record to the file
func AddTransaction(transaction models.Transaction) error {
	file, err := ioutil.ReadFile(transactionDataPath)
	if err != nil {
		return err
	}

	var transactions []models.Transaction
	if err := json.Unmarshal(file, &transactions); err != nil {
		return err
	}

	// Append the new transaction
	transactions = append(transactions, transaction)

	// Marshal and save back to file
	updatedData, err := json.Marshal(transactions)
	if err != nil {
		return err
	}

	return os.WriteFile(transactionDataPath, updatedData, 0644)
}
