package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alwafauzan/merchant-bank-api/internal/auth"
	"github.com/alwafauzan/merchant-bank-api/internal/customer"
	"github.com/alwafauzan/merchant-bank-api/internal/history"
	"github.com/alwafauzan/merchant-bank-api/internal/payment"
)

// Token blacklist to store invalidated tokens
var tokenBlacklist = make(map[string]bool)

type Server struct {
	customers []customer.Customer
}

func NewServer(customers []customer.Customer) *Server {
	return &Server{customers: customers}
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds customer.Customer
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, authenticated := customer.Authenticate(s.customers, creds.Username, creds.Password)
	if !authenticated {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	history.LogActivity("data/history.json", fmt.Sprintf("User   %s logged in", user.Username))
	w.Write([]byte(token))
}

func (s *Server) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Check if token is blacklisted
	if tokenBlacklist[token] {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var pay payment.Payment
	err = json.NewDecoder(r.Body).Decode(&pay)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	registeredIDs := make([]string, len(s.customers))
	for i, cust := range s.customers {
		registeredIDs[i] = cust.ID
	}

	err = payment.ProcessPayment(pay.FromCustomerID, pay.ToCustomerID, pay.Amount, registeredIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history.LogActivity("data/history.json", fmt.Sprintf("User   %s made a payment to %s", claims.Username, pay.ToCustomerID))
	w.Write([]byte("Payment processed"))
}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Add token to blacklist
	tokenBlacklist[token] = true

	history.LogActivity("data/history.json", fmt.Sprintf("User   %s logged out", claims.Username))
	w.Write([]byte("Logged out"))
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}