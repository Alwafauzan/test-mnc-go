package merchant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Merchant represents a merchant entity.
type Merchant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MerchantService defines the methods available for the merchant service.
type MerchantService interface {
	LoadMerchants(filePath string) ([]Merchant, error)
}

type merchantService struct{}

// NewMerchantService returns a new instance of MerchantService.
func NewMerchantService() MerchantService {
	return &merchantService{}
}

func (s *merchantService) LoadMerchants(filePath string) ([]Merchant, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read merchants file: %v", err)
		return nil, err
	}

	var merchants []Merchant
	if err := json.Unmarshal(data, &merchants); err != nil {
		log.Printf("Failed to unmarshal merchants: %v", err)
		return nil, err
	}

	fmt.Printf("Loaded merchants: %v\n", merchants)
	return merchants, nil
}