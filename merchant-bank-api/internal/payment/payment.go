package payment

import (
	"fmt"
)

type Payment struct {
	FromCustomerID string  `json:"from_customer_id"`
	ToCustomerID   string  `json:"to_customer_id"`
	Amount         float64 `json:"amount"`
}

func ProcessPayment(fromCustomerID, toCustomerID string, amount float64, registeredCustomers []string) error {
	if !isCustomerRegistered(toCustomerID, registeredCustomers) {
		return fmt.Errorf("recipient customer is not registered")
	}

	// Here you would handle the payment logic, such as updating balances.
	// For this example, we'll just log the transaction.
	fmt.Printf("Processed payment of %.2f from %s to %s\n", amount, fromCustomerID, toCustomerID)
	return nil
}

func isCustomerRegistered(customerID string, registeredCustomers []string) bool {
	for _, id := range registeredCustomers {
		if id == customerID {
			return true
		}
	}
	return false
}