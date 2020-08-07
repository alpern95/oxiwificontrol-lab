package groupe

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/db"
	"log"
	"strconv"
)

type GroupeController struct {
}

func (controller GroupeController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/groupe").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listBornesGroupe))
	//ws.Route(ws.GET("/{borneId}").Filter(auth.BearerAuth).To(getBorne))
	ws.Route(ws.PUT("/refresh/{borneId}").Filter(auth.BearerAuth).To(refreshBorne))
        ws.Route(ws.PUT("/stop/{borneId}").Filter(auth.BearerAuth).To(stopBorne))
        ws.Route(ws.PUT("/start/{borneId}").Filter(auth.BearerAuth).To(startBorne))
    //log.Printf("BorneId: %s", ws)
	return ws
}

func listBornesGroupe(req *restful.Request, resp *restful.Response) {
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

func refreshBorne(req *restful.Request, resp *restful.Response) {
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
	log.Printf("Refresh BorneId Normale: %s", err)
	resp.WriteEntity(borne)
}

func stopBorne(req *restful.Request, resp *restful.Response) {
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
        log.Printf("Stop BorneId Normale: %s", err)
        resp.WriteEntity(borne)
}

func startBorne(req *restful.Request, resp *restful.Response) {
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
        log.Printf("Start BorneId Normale: %s", err)
        resp.WriteEntity(borne)
}
