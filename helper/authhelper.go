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

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/models"
	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenManager struct {
	D          time.Duration
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	User       *models.User
}

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
func SavePrivateToDisk(filename string, key *ecdsa.PrivateKey) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the private key to the file
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if err != nil {
		return err
	}
	return nil
}
func SavePublicToDisk(filename string, pubkey *ecdsa.PublicKey) error {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Save the public key to the file
	keyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}
	err = pem.Encode(outFile, &pem.Block{Type: "PUBLIC KEY", Bytes: keyBytes})
	if err != nil {
		return err
	}
	return nil
}
func LoadPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	// Load the private key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		return nil, err
	}
	data, _ := pem.Decode(buffer)
	key, err := x509.ParseECPrivateKey(data.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return key, nil
}
func LoadPublicKey(filename string) (*ecdsa.PublicKey, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	// Load the public key from the file
	fileInfo, _ := inFile.Stat()
	buffer := make([]byte, fileInfo.Size())
	_, err = inFile.Read(buffer)
	if err != nil {
		return nil, err
	}
	data, _ := pem.Decode(buffer)
	key, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PublicKey), nil
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
func CheckKeys(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func (t *TokenManager) GenerateToken() (string, error) {

	// Create the JWT claims
	claims := jwt.MapClaims{
		"id":  t.User.ID,
		"exp": time.Now().Add(t.D).Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(t.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t *TokenManager) VerifyToken(tokenString string) (bool, error) {
	// Parse the token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.PublicKey, nil
	})

	// Check if the token is valid
	if !token.Valid {
		return false, nil
	}
	return true, nil
}

type LoginManager struct {
	Userlogin         models.UserLogin
	HashPassword      string
	LoginErrorMessage error
	Tm                *TokenManager
	Db                database.Database
}

func (l *LoginManager) CheckPassword() (*bool, error) {

	iscorrect, err := ComparePassword(l.HashPassword, l.Userlogin.Password)
	if err != nil || !iscorrect {
		log.Default().Println(err)
		err = l.LoginErrorMessage
		return nil, err
	}
	return &iscorrect, nil

}
func (l *LoginManager) CheckUser() (*models.User, error) {
	user, err := database.Find(l.Db, l.Userlogin.Email)
	if err != nil {
		log.Default().Println(err)
		return nil, l.LoginErrorMessage
	}
	return user, nil
}
