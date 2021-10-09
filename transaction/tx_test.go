package transaction

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

// Method to test transaction sign method and verify method
func TestTransaction() {
	// Create private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	// Create public key
	publicKey := privateKey.PublicKey
	// Create transaction
	transaction := NewTransaction(privateKey, publicKey, publicKey, []byte("Hello World"), Transaction_POST, "")
	// Verify transaction
	if transaction.Verify() {
		fmt.Println("Transaction verified")
	} else {
		fmt.Println("Transaction not verified")
	}
}

// Method to test StringToKey and KeyToString
func TestKey() {
	// Create private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	// Create public key
	publicKey := privateKey.PublicKey
	// Create string from public key
	s := KeyToString(publicKey)
	// Create public key from string
	key := StringToKey(s)
	// Verify public key
	if key.X.Cmp(publicKey.X) == 0 && key.Y.Cmp(publicKey.Y) == 0 {
		fmt.Println("Key verified")
	} else {
		fmt.Println("Key not verified")
	}
}
