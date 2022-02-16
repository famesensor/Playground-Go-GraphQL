package models

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Status string `json:"status"`
}
