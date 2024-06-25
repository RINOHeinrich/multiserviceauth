package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/config"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var Loginmanager helper.LoginManager
var Tokenmanager helper.TokenManager
var Keymanager helper.KeyManager

func Login(w *http.ResponseWriter, r *http.Request) {
	userlogin := models.UserLogin{}
	json.NewDecoder(r.Body).Decode(&userlogin)
	Tokenmanager.LoadConfig(&config.Config)
	Loginmanager = helper.LoginManager{
		Userlogin:         userlogin,
		HashPassword:      "",
		LoginErrorMessage: errors.New("invalid username or password"),
		Tm:                &Tokenmanager,
		Db:                &DB,
	}
	user, err := Loginmanager.CheckUser()
	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	Loginmanager.HashPassword = user.Password
	_, err = Loginmanager.CheckPassword()
	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	Keymanager.LoadConfig(&config.Config.Keyconfig)
	Tokenmanager.User = user
	Tokenmanager.Keymanager = Keymanager
	token, err := Tokenmanager.GenerateToken()
	if err != nil {
		log.Default().Println(err)
		return
	}

	json.NewEncoder(*w).Encode(token)
}
