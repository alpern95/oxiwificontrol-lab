package credential

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/db"
	//"log"
)

type CredentialController struct {
}

func (controller CredentialController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/credential").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	//ws.Route(ws.POST("/").Filter(auth.BearerAuth).To(createCredential))
    ws.Route(ws.POST("/").To(createCredential))
    
	//ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listCredentials))
	ws.Route(ws.GET("/").To(listCredentials))
	
	ws.Route(ws.GET("/{credentialId}").Filter(auth.BearerAuth).To(getCredential))

	ws.Route(ws.PUT("/{credentialId}").Filter(auth.BearerAuth).To(updateCredential))
	
	//ws.Route(ws.DELETE("/{credentialId}").Filter(auth.BearerAuth).To(deleteCredential))
	ws.Route(ws.DELETE("/{credentialId}").To(deleteCredential))

	return ws
}

func createCredential(req *restful.Request, resp *restful.Response) {
	credential := Credential{}
	err := req.ReadEntity(&credential)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	credential.ID = bson.NewObjectId()
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("credential")

	err = c.Insert(credential)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(credential)
}

func listCredentials(req *restful.Request, resp *restful.Response) {
	allCredentials := make([]Credential, 0)
	//totalCredential := make([]Credential, 0)
	totalCredential := len(allCredentials)
	//allCredentials = allCredentials,totalCredential
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("credential")
	err := c.Find(bson.M{}).All(&allCredentials)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

    resp.AddHeader("X-TOTAL-COUNT",string(totalCredential) )
	resp.WriteEntity(allCredentials)
	//resp.WriteJson("nombre" ,totalCredential)
	//resp.WriteJson(totalResp,"")
	
}

func getCredential(req *restful.Request, resp *restful.Response) {
	credentialId := req.PathParameter("credentialId")
	credential := Credential{}
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("credential")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(credentialId)}).One(&credential)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(credential)
}

func updateCredential(req *restful.Request, resp *restful.Response) {
	credentialId := req.PathParameter("credentialId")
	credential := Credential{}
	err := req.ReadEntity(&credential)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}

	credential.ID = bson.ObjectIdHex(credentialId)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("credential")
	err = c.Update(bson.M{"_id": credential.ID}, credential)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(credential)
}

func deleteCredential(req *restful.Request, resp *restful.Response) {
	credentialId := req.PathParameter("credentialId")
	var id bson.ObjectId = bson.ObjectIdHex(credentialId) // correction bug
    //log.Printf("la credential id %s",id)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("credential")
	//err := c.RemoveId(credentialId)              // Bug 
	err := c.Remove(bson.M{"_id": &id})       // correction
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}
	resp.WriteHeader(200)
}
