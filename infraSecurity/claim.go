package infraSecurity

import "github.com/dgrijalva/jwt-go"

type ApiClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}