package main

import (
	"log"
	"net/http"

	"github.com/alwafauzan/merchant-bank-api/internal/customer"
	"github.com/alwafauzan/merchant-bank-api/internal/server"
)

func main() {
	// Load customers from JSON file
	customers, err := customer.LoadCustomers("data/customers.json")
	if err != nil {
		log.Fatalf("Failed to load customers: %v", err)
	}

	// Create a new server with loaded customers
	srv := server.NewServer(customers)

	// Register HTTP handlers
	http.HandleFunc("/login", srv.LoginHandler)
	http.HandleFunc("/payment", srv.PaymentHandler)
	http.HandleFunc("/logout", srv.LogoutHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}