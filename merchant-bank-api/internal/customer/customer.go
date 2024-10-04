package customer

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Customer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadCustomers(filePath string) ([]Customer, error) {
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

	return customers, nil
}

func Authenticate(customers []Customer, username, password string) (*Customer, bool) {
	for _, customer := range customers {
		if customer.Username == username && customer.Password == password {
			return &customer, true
		}
	}
	return nil, false
}