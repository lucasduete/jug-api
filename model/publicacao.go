package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Publication struct {
	ID         bson.ObjectId `bson:"_id" json:id`
	Titulo     string        `bson:"titulo" json:"titulo"`
	Conteudo   string        `bson:"conteudo" json:"conteudo"`
	Data       time.Time     `bson:"data" json:"data"`
	EmailUser  string        `bson:"emailUser" json:"-"`
	Tecnologia string        `bson:"tecnologia" json:"tecnologia"`
}
