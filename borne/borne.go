package borne
import (
	"github.com/globalsign/mgo/bson" 
	//"github.com/alpern95/go-restful-api/credential"
)

type Borne struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Nom          string        `json:"nom" bson:"nom"`
	Description  string        `json:"description" bson:"description"`
	Credentials  []credential  `bson:"credentials"`
}
