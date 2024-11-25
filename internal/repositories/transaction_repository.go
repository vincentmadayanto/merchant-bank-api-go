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

func AddTransaction(transaction models.Transaction) error {
	file, err := ioutil.ReadFile(transactionDataPath)
	if err != nil {
		return err
	}

	var transactions []models.Transaction
	if err := json.Unmarshal(file, &transactions); err != nil {
		return err
	}

	transactions = append(transactions, transaction)

	updatedData, err := json.Marshal(transactions)
	if err != nil {
		return err
	}

	return os.WriteFile(transactionDataPath, updatedData, 0644)
}
