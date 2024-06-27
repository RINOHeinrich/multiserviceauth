package models

import "time"

type Dbconfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}
type Corsconfig struct {
	AccessControlAllowOrigin      string
	AccessControlAllowMethods     string
	AccessControlAllowHeaders     string
	AccessControlAllowCredentials string
}
type Keyconfig struct {
	PrivateKeyPath string
	PublicKeyPath  string
}
type Tokenconfig struct {
	Duration time.Duration
}
type Bcryptconfig struct {
	Cost int
}
