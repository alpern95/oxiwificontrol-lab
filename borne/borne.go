package borne
import (
	"github.com/globalsign/mgo/bson" )

type Borne struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Nom          string        `json:"nom" bson:"nom"`
	Description  string        `json:"description" bson:"description"`
	Groupe       string        `json:"groupe" bson:"groupe"`
}
