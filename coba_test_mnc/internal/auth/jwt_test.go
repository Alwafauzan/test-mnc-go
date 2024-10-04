package auth

import (
	"testing"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	username := "testuser"
	token, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %v, got %v", username, claims.Username)
	}
}