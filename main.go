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
    
	host := "192.168.112.10:3000"
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
