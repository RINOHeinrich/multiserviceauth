package helper

import (
	"time"

	"github.com/RINOHeinrich/multiserviceauth/models"
	jwt "github.com/golang-jwt/jwt"
)

type TokenManager struct {
	Duration   time.Duration
	Keymanager KeyManager
	User       *models.User
}

func (t *TokenManager) GenerateToken() (string, error) {

	// Create the JWT claims
	claims := jwt.MapClaims{
		"data": t.User,
		"exp":  time.Now().Add(t.Duration).Unix(),
	}

	// Create the JWT tokenø
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(t.Keymanager.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t *TokenManager) VerifyToken(tokenString string) (bool, error) {
	// Parse the token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.Keymanager.PublicKey, nil
	})

	// Check if the token is valid
	if !token.Valid {
		return false, nil
	}
	return true, nil
}
