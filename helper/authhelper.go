package helper

import (
	"errors"

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginManager struct {
	Config       models.LoginmanagerConfig
	Userlogin    models.UserLogin
	HashPassword string
	Tm           *TokenManager
	Bh           *BcryptHandler
}

func (l *LoginManager) CheckPassword() error {
	err := l.Bh.ComparePassword(l.HashPassword, l.Userlogin.Password)
	if err != nil {
		if l.Config.LoginErrorMessage == "" {
			err = errors.New("incorrect username or password")
		} else {
			err = errors.New(l.Config.LoginErrorMessage)
		}
		return err
	}
	return nil
}
func (l *LoginManager) CheckUser(Db database.Database) (models.User, error) {
	user := &models.User{}
	user, err := database.Find(Db, l.Userlogin.Login)
	user1 := models.User{}
	if err != nil {
		if l.Config.LoginErrorMessage == "" {
			err = errors.New("incorrect username or password")
		} else {
			err = errors.New(l.Config.LoginErrorMessage)
		}
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
		return ""
	}
	return string(hash)
}
func (b *BcryptHandler) ComparePassword(hashedPassword string, password string) error {
	// Compare the hashed password with the password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		//log.Println("Erreur de comparaison", err)
		return err
	}
	return nil
}
