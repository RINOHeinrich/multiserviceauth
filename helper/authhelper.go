package helper

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	// Hash the password with the default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
func ComparePassword(hashedPassword string, password string) (bool, error) {
	// Compare the hashed password with the password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
func SavePrivateToDisk(filename string, key *ecdsa.PrivateKey) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the private key to the file
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if err != nil {
		log.Fatal(err)
	}
}
func SavePublicToDisk(filename string, pubkey *ecdsa.PublicKey) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the public key to the file
	keyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		log.Fatal(err)
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PUBLIC KEY", Bytes: keyBytes})
	if err != nil {
		log.Fatal(err)
	}
}
func LoadPrivateKey(filename string) *ecdsa.PrivateKey {
	inFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	// Load the private key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := pem.Decode(buffer)
	key, err := x509.ParseECPrivateKey(data.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return key
}
func LoadPublicKey(filename string) *ecdsa.PublicKey {
	inFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	// Load the public key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := pem.Decode(buffer)
	key, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return key.(*ecdsa.PublicKey)
}
func GenerateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	// Generate the private key

	Keys, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// Get the public key
	publicKey := &Keys.PublicKey

	return Keys, publicKey
}

func GenerateToken(d time.Duration, privateKey *ecdsa.PrivateKey) (string, error) {
	// Create the JWT claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(d).Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
