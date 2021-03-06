package borne

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/oxiwificontrol-lab/auth"
	//"github.com/alpern95/go-restful-api/db"
	"github.com/alpern95/oxiwificontrol-lab/db"
	"log"
	"strconv"
)

type BorneController struct {
}

func (controller BorneController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/borne").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/").Filter(auth.BearerAuth).To(createBorne))
    //ws.Route(ws.POST("/").To(createBorne))

	ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listBornes))
	//ws.Route(ws.GET("/").To(listBornes))

	ws.Route(ws.GET("/{borneId}").Filter(auth.BearerAuth).To(getBorne))
	//ws.Route(ws.GET("/{borneId}").To(getBorne))

	// Creation nouvelle route 30/07/2020
	ws.Route(ws.GET("/groupe/{groupe}").Filter(auth.BearerAuth).To(getBorneGroupe))
	
	//ws.Route(ws.GET("/{utilisateur}").Filter(auth.BearerAuth).To(getBorneuser))

	ws.Route(ws.PUT("/{borneId}").Filter(auth.BearerAuth).To(updateBorne))
	//ws.Route(ws.PUT("/{borneId}").To(updateBorne))
    
	ws.Route(ws.DELETE("/{borneId}").Filter(auth.BearerAuth).To(deleteBorne))
	//ws.Route(ws.DELETE("/{borneId}").To(deleteBorne))

    //log.Printf("BorneId: %s", ws)
	return ws
}

func createBorne(req *restful.Request, resp *restful.Response) {
	borne := Borne{}
	err := req.ReadEntity(&borne)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	borne.ID = bson.NewObjectId()
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")

	err = c.Insert(borne)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(borne)
}

func listBornes(req *restful.Request, resp *restful.Response) {
	allBornes := make([]Borne, 0)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")
	err := c.Find(bson.M{}).All(&allBornes)
	if err != nil {
		resp.WriteError(500, err)
		//log.Printf("BorneId: %s", err)
		return
	}
	//log.Printf("BorneId OK: %s", err)
    totalBorne := len(allBornes)
    //log.Printf("talaborne: %s", totalBorne)
    resp.AddHeader("X-TOTAL-COUNT", strconv.Itoa(totalBorne) )
    resp.AddHeader("Content-Range", strconv.Itoa(totalBorne) )
    resp.AddHeader("Access-Control-Expose-Headers", "X-Total-Count" ) 
    ////resp.AddHeader("Access-Control-Allow-Origin","http://192.168.112.10:3001")
    //resp.AddHeader("Access-Control-Allow-Origin","http://192.168.1.32:3001")
	resp.WriteEntity(allBornes)
}

func getBorne(req *restful.Request, resp *restful.Response) {
	borneId := req.PathParameter("borneId")
	borne := Borne{}
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(borneId)}).One(&borne)
	if err != nil {
		if err == mgo.ErrNotFound {
		    log.Printf("BorneId Mongo: %s", err)
			resp.WriteError(404, err)
		} else {
		    log.Printf("BorneId: %s", err)
			resp.WriteError(500, err)
		}
		return
	}
	//log.Printf("BorneId Normale: %s", err)
	resp.WriteEntity(borne)
}

func getBorneGroupe(req *restful.Request, resp *restful.Response) {
       //utilisateur := req.PathParameter("utilisateur")
        borne := Borne{}
        session := db.NewDBSession()
        defer session.Close()
        resp.WriteEntity(borne)
}


func updateBorne(req *restful.Request, resp *restful.Response) {
	borneId := req.PathParameter("borneId")
	borne := Borne{}
	err := req.ReadEntity(&borne)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}

	borne.ID = bson.ObjectIdHex(borneId)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")
	err = c.Update(bson.M{"_id": borne.ID}, borne)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(borne)
}

func deleteBorne(req *restful.Request, resp *restful.Response) {
	borneId := req.PathParameter("borneId")
	var id bson.ObjectId = bson.ObjectIdHex(borneId) // correction bug
    //log.Printf("la borne id %s",id)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")
	//err := c.RemoveId(borneId)              // Bug 
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
