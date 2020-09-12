package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
    "github.com/alpern95/oxiwificontrol-lab/groupe"
    "github.com/alpern95/oxiwificontrol-lab/borne"
    "github.com/alpern95/oxiwificontrol-lab/auth"
	"github.com/alpern95/oxiwificontrol-lab/db"
	"github.com/joho/godotenv"
	"os"
	"os/user"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db.InitDB()

    //Add user path
	//usr, err := user.Current()

	wsContainer := restful.NewContainer()
	wsContainer.Add(borne.BorneController{}.AddRouters())
	wsContainer.Add(groupe.GroupeController{}.AddRouters())
	wsContainer.Add(auth.UserController{}.AddRouters())

	// Add container filter to enable CORS

	cors := restful.CrossOriginResourceSharing{
	    //Debug: true,
        //AllowedMethods: []string{"GET", "POST", "PUT","PATCH", "DELETE","HEAD", "OPTIONS"},
        AllowedMethods: []string{"GET", "POST", "PUT","PATCH", "DELETE","HEAD"},
        AllowedHeaders: []string{"Content-Type","X-Total-Count", "Accept","Authorization"},
        //AllowedHeaders: []string{"Content-Type"},     
        //ExposeHeaders:  []string{"X-Total-Count","x-custom-header","Access-Control-Allow-Origin"},
        ExposeHeaders:  []string{"X-Total-Count","Access-Control-Allow-Origin","Content-Range"},
        CookiesAllowed: false,
        Container:      wsContainer}

	wsContainer.Filter(cors.Filter)
    wsContainer.Filter(wsContainer.OPTIONSFilter)
    
	//host := "192.168.112.10:3000"
	host := "127.0.0.1:8081"
	//host := goDotEnvVariable("OXIWIFICONTROLADDR")
	//port := goDotEnvVariable("OXIWIFICONTROLPORT")
	//log.Printf("listening on: %s", host+":"+port)
	//hostaddr := host+":"+port
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func goDotEnvVariable(key string) string {

  // load .env file
  usr, err := user.Current()
  //err := godotenv.Load("/home/alain/go-restful-api/.env")
  err =   godotenv.Load(usr.HomeDir+"/go-restful-api/.env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}
