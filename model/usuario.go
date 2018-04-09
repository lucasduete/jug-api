package model

type User struct {
	Nome     string `bson:nome json:"nome"`
	Username string `bson:username json:"username"`
	Email    string `bson:email json:"email"`
	Senha    string `bson:senha json:"senha"`
}
