package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	//"github.com/alpern95/go-restful-api/auth"
    "./auth"
	//"github.com/alpern95/go-restful-api/borne"
    "./borne"
	//"github.com/alpern95/go-restful-api/groupe"
	"./groupe"
	"github.com/alpern95/go-restful-api/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db.InitDB()

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
	host := goDotEnvVariable("OXIWIFICONTROLADDR")
	port := goDotEnvVariable("OXIWIFICONTROLPORT")
	//log.Printf("listening on: %s", host+":"+port)
	hostaddr := host+":"+port
	log.Printf("listening on: %s", hostaddr)
	server := &http.Server{Addr: hostaddr, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}
