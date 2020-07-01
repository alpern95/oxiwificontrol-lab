package groupe
import (
	"github.com/globalsign/mgo/bson" )

type Groupe struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Nom          string `json:"nom" bson:"nom"`
	Description  string `json:"description" bson:"description"`
}
