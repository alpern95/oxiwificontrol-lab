package auth
  import "github.com/globalsign/mgo/bson"

  type User struct {
    ID       bson.ObjectId `json:"id" bson:"_id"`
    Username string `json:"username"  bson:"username"`
    Password string `json:"password"  bson:"password"`
    Email    string `json:"email"     bson:"email"`
    Groupe   string `json:"groupe"    bson:"groupe"`
    Token    string `json:"token"     bson:"token"`
  }
