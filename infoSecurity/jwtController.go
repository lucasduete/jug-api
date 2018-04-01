package infoSecurity

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("S3CR3t$K3Y_F0R_4p1_%JwT%_JUG-4p1")

func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &ApiClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*ApiClaims)
	return token.Valid, claims.Email
}
