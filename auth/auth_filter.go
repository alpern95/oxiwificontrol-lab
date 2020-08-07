package auth

import (
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/oxiwificontrol-lab/db"
)

// BearerAuth is used by all other endpoints to performan bearer token authorization
func BearerAuth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	tokenHeader := req.HeaderParameter("Authorization")
	if len(tokenHeader) == 0 {
		resp.WriteErrorString(401, "not authorized")
		return
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		resp.WriteErrorString(401, "not authorized")
		return
	}

	// the real token
	token := splitted[1]

	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("user")
	user := User{}
	err := c.Find(bson.M{"token": token}).One(&user)

	if err != nil {
		resp.WriteErrorString(401, "Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}
