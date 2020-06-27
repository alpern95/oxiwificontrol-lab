package book
import (
	"github.com/globalsign/mgo/bson" )

type Book struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}
