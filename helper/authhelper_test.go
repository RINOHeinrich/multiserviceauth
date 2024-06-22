package helper

import (
	"fmt"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	privateKey, publicKey := GenerateKeys()

	// Check if the private key is nil
	if privateKey == nil {
		t.Errorf("Expected non-nil private key, got nil")
	}
	fmt.Println("Private Key: ", privateKey)
	// Check if the public key is nil
	if publicKey == nil {
		t.Errorf("Expected non-nil public key, got nil")
	}
	fmt.Println("Public Key: ", publicKey)
	// No need to check the type of privateKey and publicKey
}
