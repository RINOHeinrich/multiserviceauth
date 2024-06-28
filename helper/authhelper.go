package helper

import (
	"log"

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginManager struct {
	Userlogin         models.UserLogin
	HashPassword      string
	LoginErrorMessage error
	Tm                *TokenManager
	Db                database.Database
	Bh                *BcryptHandler
}

func (l *LoginManager) CheckPassword() error {

	err := l.Bh.ComparePassword(l.HashPassword, l.Userlogin.Password)
	if err != nil {
		err = l.LoginErrorMessage
		return err
	}
	return nil

}
func (l *LoginManager) CheckUser(Db database.Database) (models.User, error) {
	user := &models.User{}
	user, err := database.Find(Db, l.Userlogin.Login)
	user1 := models.User{}
	if err != nil {
		return user1, err
	}

	return *user, nil
}

type BcryptHandler struct {
	Config models.Bcryptconfig
}

func (b *BcryptHandler) HashPassword(password string) string {
	// Hash the password with the default cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erreur de hashing", err)
	}
	return string(hash)
}
func (b *BcryptHandler) ComparePassword(hashedPassword string, password string) error {
	// Compare the hashed password with the password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
