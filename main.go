package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// The JWT secret (remove '0x' prefix if present)
	secretHex := "c4d70beb372fc886335d5bef3aabd63b4324b621132f5d8de67a06ddd0405fcd"
	secret, err := hex.DecodeString(secretHex)
	if err != nil {
		log.Fatalf("Failed to decode hex secret: %v", err)
	}

	// Create the claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(), // Set expiration to 1 hour from now
		"iat": time.Now().Unix(),
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	fmt.Printf("Authorization: Bearer %s\n", tokenString)
}
