package models

type JWTOptions struct {
	SecretKey       string
	ExpiresDuration int
}
