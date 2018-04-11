package infraSecurity

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"jug-api/dao"
	redis2 "github.com/go-redis/redis"
)

var secret = []byte("S3CR3t$K3Y_F0R_4p1_%JwT%_JUG-4p1")

func GenerateToken(email string) (string, error) {

	claims := ApiClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		fmt.Println("Problema ao gerar Token")
	}

	return tokenString, nil
}

func ValidateToken(myToken string) (bool, string) {

	redis := dao.GetConnectionRedis()
	defer redis.Close()

	redisToken, err := redis.Get("myToken").Result()
	if err != redis2.Nil {
		return false, redisToken
	}

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
