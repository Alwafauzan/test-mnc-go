// package main

// import (
// 	"fmt"

// 	"github.com/alwafauzan/test-mnc-go/internal/customer"
// 	"github.com/alwafauzan/test-mnc-go/internal/history"
// 	"github.com/alwafauzan/test-mnc-go/internal/merchant"
// )

// func main() {
// 	fmt.Println("Starting the application...")

// 	// Example usage
// 	customer.LoadCustomers("data/customers.json")
// 	merchant.LoadMerchants("data/merchants.json")
// 	history.LogActivity("data/history.json", "Sample activity")
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alwafauzan/test-mnc-go/internal/customer"
	"github.com/alwafauzan/test-mnc-go/internal/history"
	"github.com/alwafauzan/test-mnc-go/internal/merchant"
)

func main() {
    fmt.Println("Starting the application...")

    http.HandleFunc("/customers", getCustomers)
    http.HandleFunc("/merchants", getMerchants)
    http.HandleFunc("/log", logActivity)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
    customerService := customer.NewCustomerService()
    customers, err := customerService.LoadCustomers("data/customers.json")
    if err != nil {
        http.Error(w, "Error loading customers", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(customers)
}

func getMerchants(w http.ResponseWriter, r *http.Request) {
    merchantService := merchant.NewMerchantService()
    merchants, err := merchantService.LoadMerchants("data/merchants.json")
    if err != nil {
        http.Error(w, "Error loading merchants", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(merchants)
}

func logActivity(w http.ResponseWriter, r *http.Request) {
    historyService := history.NewHistoryService()
    err := historyService.LogActivity("data/history.json", "Sample activity")
    if err != nil {
        http.Error(w, "Error logging activity", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Activity logged"))
}