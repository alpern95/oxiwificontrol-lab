package credential
import (
	"github.com/globalsign/mgo/bson" )

type Credential struct {
	ID             bson.ObjectId `jon:"id"             bson:"_id"`
	Device         string        `json:"device"        bson:"device"`
	Adresse        string        `json:"adresse"       bson:"adresse"`
	Username       string        `json:"username"      bson:"username"`
	Password       string        `json:"password"      bson:"password"`
	Enablepassword string        `json:"enablepassword bson:"enablepassword"`
	Interface      string        `json:"interface"     bson:"insterface"`
}
