package groupe

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/db"
)

type GroupeController struct {
}

func (controller GroupeController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/groupe").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	//ws.Route(ws.POST("/").Filter(auth.BearerAuth).To(createGroupe))
    ws.Route(ws.POST("/").To(createGroupe))
    
	ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listGroupes))
	//ws.Route(ws.GET("/").To(listGroupes))
	
	ws.Route(ws.GET("/{groupeId}").Filter(auth.BearerAuth).To(getGroupe))

	ws.Route(ws.PUT("/{groupeId}").Filter(auth.BearerAuth).To(updateGroupe))
	
	ws.Route(ws.DELETE("/{groupeId}").Filter(auth.BearerAuth).To(deleteGroupe))

	return ws
}

func createGroupe(req *restful.Request, resp *restful.Response) {
	groupe := Groupe{}
	err := req.ReadEntity(&groupe)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	groupe.ID = bson.NewObjectId()
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("groupe")

	err = c.Insert(groupe)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(groupe)
}

func listGroupes(req *restful.Request, resp *restful.Response) {
	allGroupes := make([]Groupe, 0)
	//totalGroupe := make([]Groupe, 0)
	totalGroupe := len(allGroupes)
	//allGroupes = allGroupes,totalGroupe
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("groupe")
	err := c.Find(bson.M{}).All(&allGroupes)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

    resp.AddHeader("X-TOTAL-COUNT",string(totalGroupe) )
	resp.WriteEntity(allGroupes)
	//resp.WriteJson("nombre" ,totalGroupe)
	//resp.WriteJson(totalResp,"")
	
}

func getGroupe(req *restful.Request, resp *restful.Response) {
	groupeId := req.PathParameter("groupeId")
	groupe := Groupe{}
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("groupe")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(groupeId)}).One(&groupe)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(groupe)
}

func updateGroupe(req *restful.Request, resp *restful.Response) {
	groupeId := req.PathParameter("groupeId")
	groupe := Groupe{}
	err := req.ReadEntity(&groupe)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}

	groupe.ID = bson.ObjectIdHex(groupeId)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("groupe")
	err = c.Update(bson.M{"_id": groupe.ID}, groupe)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(groupe)
}

func deleteGroupe(req *restful.Request, resp *restful.Response) {
    groupeId := req.PathParameter("groupeId")
    var id bson.ObjectId = bson.ObjectIdHex(groupeId) // correction bug
    session := db.NewDBSession()
    defer session.Close()
    c := session.DB("").C("groupe")
    //err := c.RemoveId(groupeId)
    err := c.Remove(bson.M{"_id": &id})
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
