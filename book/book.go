package book

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Book struct {
	ID          bson.ObjectId `json:"id"`
	Name        string        `json:"name"`
	Author      string        `json:"author"`
}
