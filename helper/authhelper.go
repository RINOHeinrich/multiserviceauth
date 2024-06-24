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
