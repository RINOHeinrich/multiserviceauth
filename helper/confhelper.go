package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/RINOHeinrich/multiserviceauth/models"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Dbconfig    models.Dbconfig
	Corsconfig  models.Corsconfig
	Keyconfig   models.Keyconfig
	Tokenconfig models.Tokenconfig
}

func (a *AppConfig) LoadConfig(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}
	tokenduration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRES_IN"))
	if err != nil {
		return err
	}
	a.Dbconfig.DBHost = os.Getenv("DB_HOST")
	a.Dbconfig.DBPort = port
	a.Dbconfig.DBUser = os.Getenv("DB_USER")
	a.Dbconfig.DBPassword = os.Getenv("DB_PASSWORD")
	a.Dbconfig.DBName = os.Getenv("DB_NAME")
	a.Corsconfig.AccessControlAllowOrigin = os.Getenv("CORS_ALLOW_ORIGIN")
	a.Corsconfig.AccessControlAllowMethods = os.Getenv("CORS_ALLOW_METHODS")
	a.Corsconfig.AccessControlAllowHeaders = os.Getenv("CORS_ALLOW_HEADERS")
	a.Corsconfig.AccessControlAllowCredentials = os.Getenv("CORS_ALLOW_CREDENTIALS")
	a.Keyconfig.PrivateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
	a.Keyconfig.PublicKeyPath = os.Getenv("PUBLIC_KEY_PATH")
	a.Tokenconfig.Duration = tokenduration
	return nil
}
