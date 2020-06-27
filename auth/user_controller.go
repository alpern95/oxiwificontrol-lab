package auth

import (
  //"log"
  "strings"
  "time"

  "github.com/emicklei/go-restful"
  "github.com/globalsign/mgo"
  "github.com/globalsign/mgo/bson"
  "github.com/rs/xid"
  "gitlab.com/bytecraze/go-restful-api/db"
  "golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

func (controller UserController) AddRouters() *restful.WebService {
  ws := new(restful.WebService)
  ws.Path("/api/v1/user").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
  ws.Route(ws.POST("/register").To(register))
  ws.Route(ws.POST("/login").To(login))
  return ws
}

func hashAndSalt(pwd string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
  if err != nil {
    return pwd, err
  }
  return string(hash), nil
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
  byteHash := []byte(hashedPwd)
  err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
  return err == nil
}

func register(req *restful.Request, resp *restful.Response) {
  user := &User{}
  err := req.ReadEntity(user)
  if err != nil {
    resp.WriteHeaderAndEntity(400, "invalid request")
    return
  }
  if len(strings.TrimSpace(user.Username)) < 6 {
    resp.WriteHeaderAndEntity(400, "username too short")
    return
  }
  if len(strings.TrimSpace(user.Password)) < 8 {
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
