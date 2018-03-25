package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Publication struct {
	ID         bson.ObjectId `bson:_id json:id`
	Titulo     string        `bson:titulo json:titulo`
	Conteudo   string        `bson:conteudo json:conteudo`
	Data       time.Time     `bson:data json:data`
	EmailUser  string        `bson:emailUser json:emailUser`
	tecnologia string        `bson:tecnologia json:tecnologia`
}
