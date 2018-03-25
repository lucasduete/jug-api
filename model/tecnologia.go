package model

import "gopkg.in/mgo.v2/bson"

type Tecnology struct {
	ID   bson.ObjectId `bson:_id json:id`
	Nome string        `bson:nome json:nome`
}
