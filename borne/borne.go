package borne
import (
	"github.com/globalsign/mgo/bson"
	//"time" 
	//"github.com/fxtlabs/date"
)

type Borne struct {
	ID           bson.ObjectId `json:"id"          bson:"_id"`
	Nom            string   `json:"nom"            bson:"nom"`
	Description    string   `json:"description"    bson:"description"`
	Device         string   `json:"device"         bson:"device"`
	Adresse        string   `json:"adresse"        bson:"adresse"`
	Groupe         string   `json:"groupe"         bson:"groupe"`
	Modele         string   `json:"modele"         bson:"modele"`
	Username       string   `json:"username"       bson:"username"`
	Password       string   `json:"password"       bson:"password"`
	Enablepassword string   `json:"enablepassword" bson:"enablepassword"`
	Interface      string   `json:"interface"      bson:"interface"`
	Etat           string   `json:"etat"           bson:"etat"`
	Lastrefresh    string   `json:"lastrefresh"    bson:"lastrefresh"`
}
