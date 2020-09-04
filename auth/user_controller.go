package auth

import (
	"log"
	"strings"
	"time"
    "strconv"
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/rs/xid"
	//"../db"
	"github.com/alpern95/go-restful-api/db"
	"golang.org/x/crypto/bcrypt"
	//"auth"
)

type UserController struct {
}

func (controller UserController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("user/login").To(login))
	ws.Route(ws.POST("user/register").To(register))
    //ws.Route(ws.GET("user/users").To(listUsers))
      ws.Route(ws.GET("users/").To(listUsers))
	return ws
}

func login(req *restful.Request, resp *restful.Response) {
	user := &User{}
	err := req.ReadEntity(user)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	session := db.NewDBSession()
	defer session.Close()
	existingUser := User{}
	c := session.DB("").C("user")
	err = c.Find(bson.M{"username": user.Username}).One(&existingUser)

	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteHeaderAndEntity(400, "invalid login")
			return
		} else {
			resp.WriteHeaderAndEntity(500, "server error")
			return
		}
	}

	if !comparePasswords(existingUser.Password, user.Password) {
		resp.WriteHeaderAndEntity(400, "invalid login")
		return
	}
	existingUser.Password = ""
	resp.WriteEntity(existingUser)
}

func register(req *restful.Request, resp *restful.Response) {
	user := &User{}
	err := req.ReadEntity(user)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}
	if len(strings.TrimSpace(user.Username)) < 5 {
		resp.WriteHeaderAndEntity(400, "username too short")
		return
	}
	if len(strings.TrimSpace(user.Password)) < 5 {
		resp.WriteHeaderAndEntity(400, "password too short")
		return
	}

	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("user")
	existingUser := User{}
	err = c.Find(bson.M{"username": user.Username}).One(&existingUser)

	if err != nil {
		if err == mgo.ErrNotFound {
			hashedPwd, err := hashAndSalt(user.Password)
			if err != nil {
				resp.WriteHeaderAndEntity(500, "server error")
				return
			}
			user.ID = bson.NewObjectId()
			user.Password = hashedPwd
			user.Token = xid.NewWithTime(time.Now()).String()
			err = c.Insert(user)
			if err != nil {
				resp.WriteHeaderAndEntity(500, "server error")
				return
			}
		} else {
			resp.WriteHeaderAndEntity(500, "server error")
			return
		}
	} else {
		resp.WriteHeaderAndEntity(409, "duplicate username")
		return
	}
	user.Password = ""
	resp.WriteEntity(user)
}

// hashAndSalt takes a plain password and returns the hash of it
func hashAndSalt(pwd string) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Printf("ERROR error hashing password: %v\n", err)
		return pwd, err
	}
	return string(hash), nil
}

// comparePasswords compares the given password hash with the given plain password and
// tells the caller whether they match or not
func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	return err == nil
}

func listUsers(req *restful.Request, resp *restful.Response) {
        allUsers := make([]User, 0)
        session := db.NewDBSession()
        defer session.Close()
        c := session.DB("").C("user")
        err := c.Find(bson.M{}).All(&allUsers)
        if err != nil {
                resp.WriteError(500, err)
                //log.Printf("BorneId: %s", err)
                return
        }
    totalUser := len(allUsers)
    resp.AddHeader("X-TOTAL-COUNT", strconv.Itoa(totalUser) )
    resp.AddHeader("Content-Range", strconv.Itoa(totalUser) )
    resp.AddHeader("Access-Control-Expose-Headers", "X-Total-Count" )
    resp.WriteEntity(allUsers)
}
