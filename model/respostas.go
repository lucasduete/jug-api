package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Response struct {
	ID            bson.ObjectId `bson:_id json:"id"`
	Conteudo      string        `bson:conteudo json:"conteudo"`
	Data          time.Time     `bson:data json:"data"`
	EmailUser     string        `bson:emailUser json:"-"`
	IdPublication bson.ObjectId `bson:idPublication json:"idPublication"`
}
