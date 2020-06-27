package db

import (
	"github.com/globalsign/mgo"
)

var gSession *mgo.Session = nil

func InitDB() error {
	session, err := mgo.Dial("mongodb://localhost:27017/apitest")
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	gSession = session
	return nil
}

func NewDBSession() *mgo.Session {
	return gSession.Copy()
}
