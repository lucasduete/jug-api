package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:_id json:id`
	Nome     string        `bson:nome json:nome`
	Username string        `bson:username json:username`
	Email    string        `bson:email json:email`
	Senha    string        `bson:senha json:senha`
}
