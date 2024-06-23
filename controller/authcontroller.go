package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"log"

	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var Loginmanager helper.LoginManager
var Tokenmanager helper.TokenManager

func Login(w *http.ResponseWriter, r *http.Request) {
	userlogin := models.UserLogin{}
	json.NewDecoder(r.Body).Decode(&userlogin)
	Tokenmanager = helper.TokenManager{
		D: 1 * time.Hour,
	}
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
	Tokenmanager.User = user
	privatekeypath := "config/keys/private.pem"
	publickeypath := "config/keys/public.pem"
	if !helper.CheckKeys(privatekeypath) {
		privatekey, pubkey := helper.GenerateKeys()
		helper.SavePublicToDisk(publickeypath, pubkey)
		helper.SavePrivateToDisk(privatekeypath, privatekey)
	}
	Tokenmanager.PrivateKey, err = helper.LoadPrivateKey(privatekeypath)
	if err != nil {
		log.Default().Println(err)
		return
	}
	Tokenmanager.PublicKey, err = helper.LoadPublicKey(publickeypath)
	if err != nil {
		log.Default().Println(err)
		return
	}

	token, err := Tokenmanager.GenerateToken()
	if err != nil {
		log.Default().Println(err)
		return
	}
	json.NewEncoder(*w).Encode(token)
}
