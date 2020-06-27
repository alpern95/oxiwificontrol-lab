package auth

type User struct {
  Username string `json:"username" bson:"username"`
  Password string `json:"password" bson:"password"`
  Token    string `json:"token" bson:"token"`
}
