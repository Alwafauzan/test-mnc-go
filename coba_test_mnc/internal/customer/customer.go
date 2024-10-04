package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Customer represents a customer entity.
type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CustomerService defines the methods available for the customer service.
type CustomerService interface {
	LoadCustomers(filePath string) ([]Customer, error)
}

type customerService struct{}

// NewCustomerService returns a new instance of CustomerService.
func NewCustomerService() CustomerService {
	return &customerService{}
}

func (s *customerService) LoadCustomers(filePath string) ([]Customer, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read customers file: %v", err)
		return nil, err
	}

	var customers []Customer
	if err := json.Unmarshal(data, &customers); err != nil {
		log.Printf("Failed to unmarshal customers: %v", err)
		return nil, err
	}

	fmt.Printf("Loaded customers: %v\n", customers)
	return customers, nil
}