package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/alpern95/go-restful-api/auth"
	"github.com/alpern95/go-restful-api/borne"
	//"github.com/alpern95/go-restful-api/groupe"
	"github.com/alpern95/go-restful-api/db"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db.InitDB()

	wsContainer := restful.NewContainer()
	wsContainer.Add(borne.BorneController{}.AddRouters())
	wsContainer.Add(auth.UserController{}.AddRouters())

	// Add container filter to enable CORS

	cors := restful.CrossOriginResourceSharing{
        //AllowedDomains: []string{"http://192.168.1.32:3001"}, 
        //AllowedHeaders: []string{"Origin","Content-Type", "Accept", "Authorization"},
        AllowedMethods: []string{"GET", "POST", "PUT","PATCH", "DELETE","HEAD", "OPTIONS"},
        ExposeHeaders:  []string{"X-Total-Count","x-custom-header","Access-Control-Allow-Origin"},
        CookiesAllowed: false,
        Container:      wsContainer}

	wsContainer.Filter(cors.Filter)

	host := "192.168.1.32:3000"
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
