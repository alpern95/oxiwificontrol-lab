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
	    //Debug: true,
        AllowedMethods: []string{"GET", "POST", "PUT","PATCH", "DELETE","HEAD", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type","X-Total-Count", "Accept"},
        //AllowedHeaders: []string{"Content-Type"},     
        ExposeHeaders:  []string{"X-Total-Count","x-custom-header","Access-Control-Allow-Origin"},
        //ExposeHeaders:  []string{"X-Total-Count","Access-Control-Allow-Origin"},
        CookiesAllowed: false,
        Container:      wsContainer}

	wsContainer.Filter(cors.Filter)
    wsContainer.Filter(wsContainer.OPTIONSFilter)
    
	host := "192.168.1.32:3000"
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
