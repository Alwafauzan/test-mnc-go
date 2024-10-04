package config

import (
	"os"
)

func GetJWTSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "default_secret_key"
	}
	return secretKey
}