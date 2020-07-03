package borne

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/db"
	"log"
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

	ws.Route(ws.PUT("/{borneId}").Filter(auth.BearerAuth).To(updateBorne))
	
	//ws.Route(ws.DELETE("/{borneId}").Filter(auth.BearerAuth).To(deleteBorne))
	ws.Route(ws.DELETE("/{borneId}").To(deleteBorne))

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
	//totalBorne := make([]Borne, 0)
	totalBorne := len(allBornes)
	//allBornes = allBornes,totalBorne
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("borne")
	err := c.Find(bson.M{}).All(&allBornes)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

    resp.AddHeader("X-TOTAL-COUNT",string(totalBorne) )
	resp.WriteEntity(allBornes)
	//resp.WriteJson("nombre" ,totalBorne)
	//resp.WriteJson(totalResp,"")
	
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
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

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
