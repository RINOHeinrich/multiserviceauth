package middleware

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type CORSHandler struct {
	AccessControlAllowOrigin      string
	AccessControlAllowMethods     string
	AccessControlAllowHeaders     string
	AccessControlAllowCredentials string
	IgnoredOptions                bool
}

func (c *CORSHandler) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the headers for CORS
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Origin", c.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", c.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", c.AccessControlAllowHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", c.AccessControlAllowCredentials)
		// If this is a preflight request, we don't want to continue
		if r.Method == "OPTIONS" && c.IgnoredOptions {
			return
		}
		next(w, r)
	}
}
func (c *CORSHandler) LoadConfig(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return err
	}
	c.AccessControlAllowOrigin = os.Getenv("CORS_ALLOW_ORIGIN")
	c.AccessControlAllowMethods = os.Getenv("CORS_ALLOW_METHODS")
	c.AccessControlAllowHeaders = os.Getenv("CORS_ALLOW_HEADERS")
	c.AccessControlAllowCredentials = os.Getenv("CORS_ALLOW_CREDENTIALS")

	return nil
}
