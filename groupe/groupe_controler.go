package groupe

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/db"
	"log"
	"strconv"
	"../oxiwificontrolssh"
	"fmt"  // pour debug 
	"reflect"  //pour debug
	"time"
	"strings"
)

const (
        HUAWEI = "huawei"
        H3C    = "h3c"
        CISCO  = "cisco"
        LINUX  = "linux"
        EXOS   = "exos"
)

type GroupeController struct {
}

type User struct {
    ID       bson.ObjectId `json:"id" bson:"_id"`
    Username string `json:"username"  bson:"username"`
    Password string `json:"password"  bson:"password"`
    Email    string `json:"email"     bson:"email"`
    Role     string `json:"role"      bson:"role"`
    Token    string `json:"token"     bson:"token"`
  }


func (controller GroupeController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/groupe").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

    //test deuxieme premier ok
	ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listBornesGroupe))
	//ws.Route(ws.GET("/{groupe}").Filter(auth.BearerAuth).To(listBornesGroupe))

	ws.Route(ws.PUT("/refresh/{borneId}").Filter(auth.BearerAuth).To(refreshBorne))
    ws.Route(ws.PUT("/stop/{borneId}").Filter(auth.BearerAuth).To(stopBorne))
    ws.Route(ws.PUT("/start/{borneId}").Filter(auth.BearerAuth).To(startBorne))
    //log.Printf("BorneId: %s", ws)
	return ws
}

func listBornesGroupe(req *restful.Request, resp *restful.Response) {

    // Determiner le token
    Bearer := req.Request.Header["Authorization"]
    var token string
    token = Bearer[0]
    if strings.HasPrefix(token, "Bearer") {
        token = strings.TrimPrefix(token, "Bearer ")
    }else {
    	return
    }
    log.Printf("le token est %s", token)

    //Find the role with the token
    session := db.NewDBSession()
    defer session.Close()
    c := session.DB("").C("user")
    result := User{}
    err := c.Find(bson.M{"token": token}).One(&result)
    if err != nil {
      if err == mgo.ErrNotFound {
          log.Printf("not found token", err)
          //return
      }else {
          log.Printf("erreur recherche token %s",err)
      }
    }
    log.Printf("Le resultat de la recherche est : %s", result)
    log.Printf("le type de result est: %s",reflect.TypeOf(result))
    log.Print("le role est : ", result.Role)


    // Lister les bornes du groupe == role
     
	allBornes := make([]Borne, 0)
	session = db.NewDBSession()
	defer session.Close()
	c = session.DB("").C("borne")
	//err = c.Find(bson.M{}).All(&allBornes)
	err = c.Find(bson.M{"groupe": result.Role }).All(&allBornes)
	//err := c.Find(bson.M{"groupe": groupe}).All(&allBornes)]]
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
	// faire un acces ssh à la borne
	cmds := make([]string, 0)
	user := borne.Username
	port := borne.Interface
	password := borne.Password
	ipPort := borne.Adresse+":22"
	brand, err := ssh.GetSSHBrand(user, password, ipPort)
    if err != nil {
    	fmt.Println("GetSSHBrand err:\n", err.Error())
    }
    fmt.Println("Device brand is: ", brand)

    if brand != HUAWEI && brand != H3C && brand != CISCO && brand != LINUX && brand != EXOS {
            //LogDebug("Detection  <brand> = ", brand)
    }
    switch brand {
    case HUAWEI:
            //session.WriteChannel(HuaweiNoPage)
            break
    case H3C:
            //session.WriteChannel(H3cNoPage)
            break
    case CISCO:
            //session.WriteChannel(CiscoNoPage)
            break
    case LINUX:
        //session.WriteChannel(LinuxNoPage)
    case EXOS:
        //run the cmds in the switch, and get the execution results
        cmds = append(cmds, "sh port "+port+" information | include "+port)
        result, err := ssh.RunCommands(user, password, ipPort, cmds...)
        if err != nil {
            fmt.Println("RunCommand err:\n", err.Error())
        }else {
        	fmt.Println("Le resultat de show port",result)

                if strings.Contains(result, "Em")  {
                    fmt.Println("Prompt true",result)
                    borne.Etat = "UP"
                }else if strings.Contains(result,"Dm") {
                	borne.Etat = "DOWN"
                }
        }
    }

   // Date Time
    maint := time.Now()
    update := maint.Format(time.RFC1123Z)
    updatetime := update

    // faire un update du champ borne status
	log.Printf("Refresh BorneId Normale at : %s", updatetime)
    borne.Lastrefresh = updatetime
    //
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
        // faire un acces ssh à la borne
        cmds := make([]string, 0)
        user := borne.Username
        password := borne.Password
        ipPort := borne.Adresse+":22"
        brand, err := ssh.GetSSHBrand(user, password, ipPort)
        if err != nil {
            fmt.Println("GetSSHBrand err:\n", err.Error())
        }
        fmt.Println("Device brand is: ", brand)

        //run the cmds in the switch, and get the execution results
        cmds = append(cmds, "uptime")
        result, err := ssh.RunCommands(user, password, ipPort, cmds...)
        if err != nil {
            fmt.Println("RunCommand err:\n", err.Error())
        }

        fmt.Println("uptime result is = : ", result)

        // Date Time
        maint := time.Now()
        update := maint.Format(time.RFC1123Z)
        updatetime := update

        // faire un update du champ borne status
        log.Printf("Refresh BorneId Normale at : %s", updatetime)
        borne.Lastrefresh = updatetime
        borne.Etat = "DOWN"
        //
        err = c.Update(bson.M{"_id": borne.ID}, borne)
        if err != nil {
                if err == mgo.ErrNotFound {
                        resp.WriteError(404, err)
                } else {
                        resp.WriteError(500, err)
                }
                return
        }
        ////

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

        ////
        // faire un acces ssh à la borne
        cmds := make([]string, 0)
        user := borne.Username
        password := borne.Password
        ipPort := borne.Adresse+":22"
        brand, err := ssh.GetSSHBrand(user, password, ipPort)
        if err != nil {
            fmt.Println("GetSSHBrand err:\n", err.Error())
        }
        fmt.Println("Device brand is: ", brand)

        //run the cmds in the switch, and get the execution results
        cmds = append(cmds, "uptime")
        result, err := ssh.RunCommands(user, password, ipPort, cmds...)
        if err != nil {
            fmt.Println("RunCommand err:\n", err.Error())
        }

        fmt.Println("uptime result is = : ", result)

        // Date Time
        maint := time.Now()
        update := maint.Format(time.RFC1123Z)
        updatetime := update

        // faire un update du champ borne status
        log.Printf("Refresh BorneId Normale at : %s", updatetime)
        borne.Lastrefresh = updatetime
        borne.Etat = "UP"
        //
        err = c.Update(bson.M{"_id": borne.ID}, borne)
        if err != nil {
                if err == mgo.ErrNotFound {
                        resp.WriteError(404, err)
                } else {
                        resp.WriteError(500, err)
                }
                return
        }
        ////

        resp.WriteEntity(borne)
}
